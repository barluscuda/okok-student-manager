package bootstrap

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/okok-student-manager/internal/config"
	"github.com/okok-student-manager/internal/handler"
	"github.com/okok-student-manager/internal/service"
	"go.uber.org/zap"
)

type App struct {
	Server *echo.Echo
	Logger *zap.Logger
}

func New(cfg config.Config) (*App, error) {
	logger, err := newLogger(cfg.Env)
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(requestLogger(logger))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		status := http.StatusInternalServerError
		if httpErr, ok := err.(*echo.HTTPError); ok {
			status = httpErr.Code
		}
		if c.Response().Committed {
			return
		}
		_ = c.JSON(status, map[string]any{
			"code":    status,
			"message": http.StatusText(status),
		})
	}

	healthService := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthService)
	e.GET("/health", healthHandler.Check)
	e.GET("/api/v1/health", healthHandler.Check)

	return &App{Server: e, Logger: logger}, nil
}

func newLogger(env string) (*zap.Logger, error) {
	if env == "production" {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}

func requestLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI: true, LogStatus: true, LogMethod: true, LogError: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			fields := []zap.Field{
				zap.String("method", values.Method),
				zap.String("uri", values.URI),
				zap.Int("status", values.Status),
			}
			if values.Error != nil {
				fields = append(fields, zap.Error(values.Error))
			}
			logger.Info("http request", fields...)
			return nil
		},
	})
}
