package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/stablecog/sc-go/cron/jobs"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/responses"
)

func (c *RestAPI) HandleGetStats(w http.ResponseWriter, r *http.Request) {
	jobRunner := jobs.JobRunner{
		Repo:   c.Repo,
		Redis:  c.Redis,
		Ctx:    context.Background(),
		Stripe: c.StripeClient,
	}
	jobRunner.GetAndSetStats(jobs.NewJobLogger("STATS"))

	res, err := c.Redis.GetGenerateUpscaleCount()
	if err != nil {
		log.Error("Error getting generate upscale count", "err", err)
		responses.ErrInternalServerError(w, r, "Unable to get stats")
		return
	}

	if err := c.Repo.WithTx(func(tx *ent.Tx) error {
		DB := tx.Client()

		res.UsersCount, res.Users24H, err = c.Repo.GetUsersCount(DB, r.Context())
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Error("Error in transaction", "err", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
