package resource

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"go.uber.org/zap"

	"datink/config"
	"datink/pkg/database"
	"datink/pkg/util/httputil"
)

type Resource struct {
	dig.In

	Config   *config.Config
	Logger   *zap.Logger
	Echo     *echo.Echo
	Database *database.DB
	Response *httputil.HttpResponse
}
