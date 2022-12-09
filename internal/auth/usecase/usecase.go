package usecase

import (
	"context"
	"go-rest-api/config"
	"go-rest-api/internal/auth"
	"go-rest-api/internal/models"
	"go-rest-api/pkg/logger"
)

const (
	basePrefix    = "api-auth:"
	cacheDuration = 3600
)

// Auth UseCase
type authUC struct {
	cfg       *config.Config
	authRepo  auth.Repository
	redisRepo auth.RedisRepository
	awsRepo   auth.AWSRepository
	logger    logger.Logger
}

func (a authUC) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	//TODO implement me
	panic("implement me")
}

// Auth UseCase constructor
func NewAuthUseCase(cfg *config.Config, authRepo auth.Repository, redisRepo auth.RedisRepository, awsRepo auth.AWSRepository, log logger.Logger) auth.UseCase {
	return &authUC{cfg: cfg, authRepo: authRepo, redisRepo: redisRepo, awsRepo: awsRepo, logger: log}
}
