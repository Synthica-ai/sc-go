package controller

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stablecog/go-apps/database"
	"github.com/stablecog/go-apps/database/repository"
	"github.com/stablecog/go-apps/server/controller/websocket"
	"github.com/stablecog/go-apps/shared"
)

type HttpController struct {
	Repo              *repository.Repository
	Redis             *database.RedisWrapper
	S3Client          *s3.Client
	S3PresignClient   *s3.PresignClient
	CogRequestUserMap *shared.SyncMap[string]
	Hub               *websocket.Hub
}
