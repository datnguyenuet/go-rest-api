package auth

import (
	"context"
	"go-rest-api/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}
