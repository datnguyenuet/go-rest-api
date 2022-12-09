package auth

import (
	"context"
	"go-rest-api/internal/models"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}
