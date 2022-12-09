package http

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"go-rest-api/config"
	"go-rest-api/internal/auth"
	"go-rest-api/internal/models"
	"go-rest-api/internal/session"
	"go-rest-api/pkg/httpErrors"
	"go-rest-api/pkg/logger"
	"go-rest-api/pkg/utils"
	"net/http"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	sessUC session.UCSession
	logger logger.Logger
}

func (a *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "auth.Register")
		defer span.Finish()

		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			utils.LogResponseError(c, a.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := a.authUC.Register(ctx, user)
		if err != nil {
			utils.LogResponseError(c, a.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		sess, err := a.sessUC.CreateSession(ctx, &models.Session{
			UserID: createdUser.User.UserID,
		}, a.cfg.Session.Expire)
		if err != nil {
			utils.LogResponseError(c, a.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		c.SetCookie(utils.CreateSessionCookie(a.cfg, sess))

		return c.JSON(http.StatusCreated, createdUser)
	}
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, sessUC session.UCSession, log logger.Logger) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC, logger: log}
}
