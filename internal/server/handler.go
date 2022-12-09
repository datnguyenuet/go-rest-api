package server

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/internal/session/usecase"
)
import (
	authRepository "go-rest-api/internal/auth/repository"
	authUseCase "go-rest-api/internal/auth/usecase"

	authHttp "go-rest-api/internal/auth/delivery/http"
	sessionRepository "go-rest-api/internal/session/repository"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	authRepo := authRepository.NewAuthRepository(s.db)
	authRedisRepo := authRepository.NewAuthRedisRepo(s.redisClient)
	authAWSRepo := authRepository.NewAuthAWSRepository(s.awsClient)
	sessRepo := sessionRepository.NewSessionRepository(s.redisClient, s.cfg)
	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.cfg, authRepo, authRedisRepo, authAWSRepo, s.logger)
	sessUC := usecase.NewSessionUseCase(sessRepo, s.cfg)
	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, sessUC, s.logger)
}
