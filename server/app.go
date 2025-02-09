package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	chiprometheus "github.com/stablecog/chi-prometheus"
	"github.com/stablecog/sc-go/cron/jobs"
	"github.com/stablecog/sc-go/database"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/analytics"
	"github.com/stablecog/sc-go/server/api/rest"
	"github.com/stablecog/sc-go/server/api/sse"
	"github.com/stablecog/sc-go/server/clip"
	"github.com/stablecog/sc-go/server/discord"
	"github.com/stablecog/sc-go/server/middleware"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/uploadapi/api"
	"github.com/stablecog/sc-go/utils"
	stripe "github.com/stripe/stripe-go/v74/client"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type AskBodySettings struct {
	Model       string `json:"model,omitempty"`
	MaxTokens   int    `json:"max_tokens,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
	TopP        int    `json:"top_p,omitempty"`
}

type AskBodyOpenAI struct {
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Model       string `json:"model,omitempty"`
	MaxTokens   int    `json:"max_tokens,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
	TopP        int    `json:"top_p,omitempty"`
	Stream      bool   `json:"stream"`
}

type AskBody struct {
	AskBodyOpenAI

	Settings AskBodySettings `json:"settings,omitempty"`
}

var Version = "dev"
var CommitMsg = "dev"

// Used to track the build time from our CI
var BuildStart = ""

func main() {
	log.Infof("SC Server: %s", Version)

	// Load .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Warn("Error loading .env file (this is fine)", "err", err)
	}

	// Custom flags
	createMockData := flag.Bool("load-mock-data", false, "Create test data in database")

	flag.Parse()

	// Setup database
	ctx := context.Background()

	// Setup sql
	log.Info("🏡 Connecting to database...")
	dbconn, err := database.GetSqlDbConn(false)
	if err != nil {
		log.Fatal("Failed to connect to database", "err", err)
		os.Exit(1)
	}
	entClient, err := database.NewEntClient(dbconn)
	if err != nil {
		log.Fatal("Failed to create ent client", "err", err)
		os.Exit(1)
	}
	defer entClient.Close()
	// Run migrations
	// We can't run on supabase, :(
	if utils.GetEnv("RUN_MIGRATIONS", "") == "true" {
		log.Info("🦋 Running migrations...")
		if err := entClient.Schema.Create(ctx); err != nil {
			log.Fatal("Failed to run migrations", "err", err)
			os.Exit(1)
		}
	}

	// Setup redis
	redis, err := database.NewRedis(ctx)
	if err != nil {
		log.Fatal("Error connecting to redis", "err", err)
		os.Exit(1)
	}

	// Setup qdrant
	// qdrantClient, err := qdrant.NewQdrantClient(ctx)
	// if err != nil {
	// 	log.Fatal("Error connecting to qdrant", "err", err)
	// 	os.Exit(1)
	// }
	// err = qdrantClient.CreateCollectionIfNotExists(false)
	// if err != nil {
	// 	log.Fatal("Error creating qdrant collection", "err", err)
	// 	os.Exit(1)
	// }

	// // Create indexes in Qdrant
	// err = qdrantClient.CreateAllIndexes()
	// if err != nil {
	// 	log.Warn("Error creating qdrant indexes", "err", err)
	// }

	// Q Throttler
	qThrottler := shared.NewQueueThrottler(ctx, redis.Client, shared.REQUEST_COG_TIMEOUT)

	// Create repository (database access)
	repo := &repository.Repository{
		DB:             entClient,
		ConnInfo:       dbconn,
		Redis:          redis,
		Ctx:            ctx,
		Qdrant:         nil,
		QueueThrottler: qThrottler,
	}

	if *createMockData {
		log.Info("🏡 Creating mock data...")
		err = repo.CreateMockData(ctx)
		if err != nil {
			log.Fatal("Failed to create mock data", "err", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Create stripe client
	stripeClient := stripe.New(utils.GetEnv("STRIPE_SECRET_KEY", ""), nil)

	app := chi.NewRouter()

	// Prometheus middleware
	promMiddleware := chiprometheus.NewMiddleware("sc-server")
	app.Use(promMiddleware)

	// Cors middleware
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins: utils.GetCorsOrigins(),
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Get models, schedulers and put in cache
	log.Info("📦 Populating cache...")
	err = repo.UpdateCache()
	if err != nil {
		// ! Not getting these is fatal and will result in crash
		panic(err)
	}
	// Update periodically
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minutes().StartAt(time.Now().Add(5 * time.Minute)).Do(func() {
		log.Info("📦 Updating cache...")
		err = repo.UpdateCache()
		if err != nil {
			log.Error("Error updating cache", "err", err)
		}
	})

	// Create SSE hub
	sseHub := sse.NewHub(redis, repo)
	go sseHub.Run()

	// Need to send keepalive every 30 seconds
	s.Every(30).Seconds().StartAt(time.Now().Add(30 * time.Second)).Do(func() {
		sseHub.BraodcastKeepalive()
	})

	// Start cron scheduler
	s.StartAsync()

	// Create analytics service
	analyticsService := analytics.NewAnalyticsService()
	defer analyticsService.Close()

	// Setup S3 Client
	region := os.Getenv("S3_IMG2IMG_REGION")
	accessKey := os.Getenv("S3_IMG2IMG_ACCESS_KEY")
	secretKey := os.Getenv("S3_IMG2IMG_SECRET_KEY")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(os.Getenv("S3_IMG2IMG_ENDPOINT")),
		Region:      aws.String(region),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	jobRunner := jobs.JobRunner{
		Repo:   repo,
		Redis:  redis,
		Ctx:    ctx,
		Meili:  database.NewMeiliSearchClient(),
		Track:  analyticsService,
		Stripe: stripeClient,
	}

	err = jobRunner.SyncMeili(jobs.NewJobLogger("MEILI_SYNC"), 1000)
	if err != nil {
		log.Fatal("Error syncing meili", "err", err)
		os.Exit(1)
	}

	s.Every(60).Seconds().Do(jobRunner.SyncMeili, jobs.NewJobLogger("MEILI_SYNC"), 100)

	s.Every(5).Minutes().Do(jobRunner.RefundOldGenerationCredits, jobs.NewJobLogger("AUTO_REFUND"))
	s.Every(60).Seconds().Do(jobRunner.AddFreeCreditsToEligibleUsers, jobs.NewJobLogger("FREE_CREDITS"))

	// Create controller
	apiTokenSmap := shared.NewSyncMap[chan requests.CogWebhookMessage]()
	hc := rest.RestAPI{
		Repo:           repo,
		Redis:          redis,
		Hub:            sseHub,
		StripeClient:   stripeClient,
		Track:          analyticsService,
		QueueThrottler: qThrottler,
		S3:             s3Client,
		Qdrant:         nil,
		Meili:          database.NewMeiliSearchClient(),
		Clip:           clip.NewClipService(redis),
		SMap:           apiTokenSmap,
	}

	// Create middleware
	mw := middleware.Middleware{
		SupabaseAuth: database.NewSupabaseAuth(),
		Repo:         repo,
		Redis:        redis,
	}

	// Create controller
	hu := api.Controller{
		Repo:  repo,
		Redis: redis,
		S3:    s3Client,
	}

	// Routes
	app.Get("/", hc.HandleHealth)
	app.Handle("/metrics", middleware.BasicAuth(promhttp.Handler(), "user", "password", "Authentication required"))
	app.Get("/clipq", hc.HandleClipQSearch)
	app.Route("/upload", func(r chi.Router) {
		// File upload
		r.Route("/", func(r chi.Router) {
			r.Get("/health", hc.HandleHealth)
			r.Route("/", func(r chi.Router) {
				r.Use(middleware.Logger)
				r.Use(mw.RateLimit(2, "srv", 1*time.Second))
				r.Use(mw.AuthMiddleware(middleware.AuthLevelAny))
				r.Post("/", hu.HandleUpload)
			})
		})
	})

	originServerURL, err := url.Parse("https://api.openai.com")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	proxy := httputil.NewSingleHostReverseProxy(originServerURL)
	proxy.ModifyResponse = hc.HandleAiChatAskResponse

	app.Route("/v2", func(r chi.Router) {
		r.Use(mw.AuthMiddleware(middleware.AuthLevelAPIToken))
		r.Use(middleware.Logger)
		r.Use(mw.RateLimit(5, "api", 1*time.Second))

		// Query credits
		r.Post("/ai-chat/api/ask", hc.HandleAiChatAsk(proxy))

		r.Post("/ai-chat/api/suggest-title", hc.HandleAiChatTitle)

		r.Patch("/user/settings", hc.HandleUpdateUserSettings)
		r.Get("/user/settings", hc.HandleGetUserSettings)

		// Create Generation
		r.Post("/user/generation", hc.HandleCreateGeneration)
	})

	// Routes that require authentication
	app.Route("/ai-chat", func(r chi.Router) {
		r.Use(mw.AuthMiddleware(middleware.AuthLevelAny))
		r.Use(middleware.Logger)

		r.Use(mw.RateLimit(10, "srv", 1*time.Second))

		// Query credits
		r.Post("/api/ask", hc.HandleAiChatAsk(proxy))

		r.Post("/api/suggest-title", hc.HandleAiChatTitle)
	})

	app.Route("/v1", func(r chi.Router) {
		r.Get("/health", hc.HandleHealth)

		// SSE
		r.Route("/sse", func(r chi.Router) {
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				sseHub.ServeSSE(w, r)
			})
		})

		// Stripe
		r.Route("/stripe", func(r chi.Router) {
			r.Use(middleware.Logger)
			r.Post("/webhook", hc.HandleStripeWebhook)
		})

		// SCWorker
		r.Route("/worker", func(r chi.Router) {
			r.Use(middleware.Logger)
			r.Post("/webhook", hc.HandleSCWorkerWebhook)
		})

		// Stats
		r.Route("/stats", func(r chi.Router) {
			r.Use(middleware.Logger)
			// 10 requests per second
			r.Use(mw.RateLimit(10, "srv", 1*time.Second))
			r.Get("/", hc.HandleGetStats)
		})

		// Gallery search
		r.Route("/gallery", func(r chi.Router) {
			r.Use(middleware.Logger)
			// 20 requests per second
			r.Use(mw.RateLimit(20, "srv", 1*time.Second))
			r.Get("/", hc.HandleQueryGallery)
		})

		// Gallery search
		r.Route("/ai_friends", func(r chi.Router) {
			r.Use(middleware.Logger)
			// 20 requests per second
			r.Use(mw.RateLimit(20, "srv", 1*time.Second))
			r.Get("/", hc.HandleGetAIFriends)
		})

		r.Route("/ai_influencers", func(r chi.Router) {
			r.Use(middleware.Logger)
			// 20 requests per second
			r.Use(mw.RateLimit(20, "srv", 1*time.Second))
			r.Get("/", hc.HandleGetAIInfluencer)
		})

		// Routes that require authentication
		r.Route("/user", func(r chi.Router) {
			r.Use(mw.AuthMiddleware(middleware.AuthLevelAny))
			r.Use(middleware.Logger)
			// 10 requests per second
			r.Use(mw.RateLimit(10, "srv", 1*time.Second))

			// Get user summary
			r.Get("/", hc.HandleGetUser)

			r.Route("/ai_voice", func(r chi.Router) {
				r.Use(middleware.Logger)
				// 20 requests per second
				r.Use(mw.RateLimit(20, "srv", 1*time.Second))
				r.Get("/", hc.HandleGetAIVoices)
				r.Post("/", hc.HandleInsertAIVoice)
				r.Patch("/", hc.HandleUpdateAIVoice)
				r.Delete("/", hc.HandleDeleteAIVoice)
			})

			r.Route("/ai_voice_settings", func(r chi.Router) {
				r.Use(middleware.Logger)
				// 20 requests per second
				r.Use(mw.RateLimit(20, "srv", 1*time.Second))
				r.Get("/", hc.HandleGetAIVoiceSettings)
				r.Post("/", hc.HandleInsertAIVoiceSettings)
				r.Patch("/", hc.HandleUpdateAIVoiceSettings)
			})

			r.Patch("/", hc.HandleUpdateUser)

			// Create Generation
			r.Post("/generation", hc.HandleCreateGeneration)
			// Mark generation for deletion
			r.Delete("/generation", hc.HandleDeleteGenerationOutputForUser)

			// Query Generation (outputs + generations)
			r.Get("/outputs", hc.HandleQueryGenerations)

			// Favorite
			r.Post("/outputs/favorite", hc.HandleFavoriteGenerationOutputsForUser)

			// Create upscale
			r.Post("/upscale", hc.HandleUpscale)

			// Query credits
			r.Get("/credits", hc.HandleQueryCredits)

			// Submit to gallery
			r.Put("/gallery", hc.HandleSubmitGenerationToGallery)

			// Subscriptions
			r.Post("/subscription/downgrade", hc.HandleSubscriptionDowngrade)
			r.Post("/subscription/checkout", hc.HandleCreateCheckoutSession)
			r.Post("/subscription/portal", hc.HandleCreatePortalSession)

			// API Tokens
			r.Post("/tokens", hc.HandleNewAPIToken)
			r.Patch("/tokens", hc.HandleUpdateAPIToken)
			r.Get("/tokens", hc.HandleGetAPITokens)
			r.Delete("/tokens", hc.HandleDeactivateAPIToken)

			// settings
			r.Patch("/settings", hc.HandleUpdateUserSettings)
			r.Get("/settings", hc.HandleGetUserSettings)

			r.Patch("/account", hc.HandleUpdateAccount)
			r.Get("/account", hc.HandleGetAccount)

			// Operations
			r.Get("/operations", hc.HandleQueryOperations)
		})

		// Admin only routes
		r.Route("/admin", func(r chi.Router) {
			r.Route("/gallery", func(r chi.Router) {
				r.Use(mw.AuthMiddleware(middleware.AuthLevelGalleryAdmin))
				r.Use(middleware.Logger)
				r.Put("/", hc.HandleReviewGallerySubmission)
			})
			r.Route("/outputs", func(r chi.Router) {
				// TODO - this is auth level gallery admin, but delete route manually enforces super admin
				r.Use(mw.AuthMiddleware(middleware.AuthLevelGalleryAdmin))
				r.Use(middleware.Logger)
				r.Delete("/", hc.HandleDeleteGenerationOutput)
				r.Get("/", hc.HandleQueryGenerationsForAdmin)
			})
			r.Route("/users", func(r chi.Router) {
				r.Use(mw.AuthMiddleware(middleware.AuthLevelSuperAdmin))
				r.Use(middleware.Logger)
				r.Get("/", hc.HandleQueryUsers)
				r.Post("/ban", hc.HandleBanUser)
			})
			r.Route("/credit", func(r chi.Router) {
				r.Use(mw.AuthMiddleware(middleware.AuthLevelSuperAdmin))
				r.Use(middleware.Logger)
				r.Get("/types", hc.HandleQueryCreditTypes)
				r.Post("/add", hc.HandleAddCreditsToUser)
			})
		})

		// Settings
		r.Route("/settings", func(r chi.Router) {
			r.Use(middleware.Logger)
			r.Use(mw.RateLimit(10, "srv", 1*time.Second))
			r.Get("/", hc.HandleGetSettings)
		})

		// Api token route
		r.Route("/generate", func(r chi.Router) {
			r.Use(mw.AuthMiddleware(middleware.AuthLevelAPIToken))
			r.Use(middleware.Logger)
			r.Use(mw.RateLimit(5, "api", 1*time.Second))
			r.Post("/", hc.HandleCreateGenerationToken)
		})

		// Api token route
		r.Route("/upload", func(r chi.Router) {
			r.Use(mw.AuthMiddleware(middleware.AuthLevelAPIToken))
			r.Use(middleware.Logger)
			r.Use(mw.RateLimit(2, "srv", 1*time.Second))
			r.Post("/", hu.HandleUpload)
		})

		// Api token route
		r.Route("/generate_similar", func(r chi.Router) {
			r.Use(mw.AuthMiddleware(middleware.AuthLevelAPIToken))
			r.Use(middleware.Logger)
			r.Use(mw.RateLimit(5, "api", 1*time.Second))
			r.Post("/", hc.HandleCreateGenerationToken)
		})

		r.Route("/upscale", func(r chi.Router) {
			r.Use(mw.AuthMiddleware(middleware.AuthLevelAPIToken))
			r.Use(middleware.Logger)
			r.Use(mw.RateLimit(5, "api", 1*time.Second))
			r.Post("/", hc.HandleCreateUpscaleToken)
		})
	})

	// This redis subscription has the following purpose:
	// After we are done processing a cog message, we want to broadcast it to
	// our subscribed SSE clients matching that stream ID
	// the purpose of this instead of just directly sending the message to the SSE is that
	// our service can scale, and we may have many instances running and we care about SSE connections
	// on all of them.
	pubsubSSEMessages := redis.Client.Subscribe(ctx, shared.REDIS_SSE_BROADCAST_CHANNEL)
	defer pubsubSSEMessages.Close()

	// Start SSE redis subscription
	go func() {
		log.Info("Listening for cog messages", "channel", shared.REDIS_SSE_BROADCAST_CHANNEL)
		for msg := range pubsubSSEMessages.Channel() {
			var sseMessage repository.TaskStatusUpdateResponse
			err := json.Unmarshal([]byte(msg.Payload), &sseMessage)
			if err != nil {
				log.Error("Error unmarshalling sse message", "err", err)
				continue
			}

			// Live page separate broadcast stream
			if sseMessage.ForLivePage {
				sseHub.BroadcastLivePageMessage(*sseMessage.LivePageMessage)
				continue
			}

			// Sanitize
			sseMessage.LivePageMessage = nil
			// The hub will broadcast this to our clients if it's supposed to
			sseHub.BroadcastStatusUpdate(sseMessage)
		}
	}()

	// This redis subscription has the following purpose:
	// For API token requests, they are synchronous with API requests
	// so we need to send the response back to the appropriate channel
	apiTokenChannel := redis.Client.Subscribe(ctx, shared.REDIS_APITOKEN_COG_CHANNEL)
	defer apiTokenChannel.Close()

	// Start SSE redis subscription
	go func() {
		log.Info("Listening for api messages", "channel", shared.REDIS_APITOKEN_COG_CHANNEL)
		for msg := range apiTokenChannel.Channel() {
			var cogMessage requests.CogWebhookMessage
			err := json.Unmarshal([]byte(msg.Payload), &cogMessage)
			if err != nil {
				log.Error("Error unmarshalling cog webhook message", "err", err)
				continue
			}

			if chl := apiTokenSmap.Get(cogMessage.Input.ID); chl != nil {
				chl <- cogMessage
			}
		}
	}()

	// Start server
	port := utils.GetEnv("PORT", "13337")
	log.Info("Starting server", "port", port)

	h2s := &http2.Server{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: h2c.NewHandler(app, h2s),
	}

	// Send discord notification
	go func() {
		err = discord.FireServerReadyWebhook(Version, CommitMsg, BuildStart)
		if err != nil {
			log.Error("Error firing discord ready webhook", "err", err)
		}
	}()
	log.Info(srv.ListenAndServe())
}
