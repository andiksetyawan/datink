package delivery

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"datink/internal/delivery/rest"
)

type Package struct {
	dig.In

	Router rest.Routes
}

func (p Package) Register(e *echo.Echo) error {
	p.Router.Register(e)
	return nil
}
