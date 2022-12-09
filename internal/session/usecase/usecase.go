package usecase

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go-rest-api/config"
	"go-rest-api/internal/models"
	"go-rest-api/internal/session"
)

type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

func (s *sessionUC) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.CreateSession")

	defer span.Finish()

	return s.sessionRepo.CreateSession(ctx, session, expire)
}

func (s *sessionUC) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.GetSessionByID")
	defer span.Finish()

	return s.sessionRepo.GetSessionByID(ctx, sessionID)
}

func (s *sessionUC) DeleteByID(ctx context.Context, sessionID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.DeleteByID")
	defer span.Finish()

	return s.sessionRepo.DeleteByID(ctx, sessionID)
}

func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.UCSession {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}
