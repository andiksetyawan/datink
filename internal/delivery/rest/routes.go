package rest

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/dig"

	_ "datink/docs"
	"datink/internal/delivery/rest/handler"
	"datink/internal/delivery/rest/middleware"
	"datink/internal/resource"
	"datink/pkg/util/httputil"
)

type Routes struct {
	dig.In

	Resource   resource.Resource
	Handler    handler.Handler
	Middleware middleware.Middlewares
}

func (r *Routes) Register(ec *echo.Echo) {
	ec.Validator = &httputil.Validator{Validator: validator.New()}

	ec.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "datink be service")
	})
	ec.GET("/docs/*", echoSwagger.WrapHandler)

	v1 := ec.Group("/v1")

	userV1 := v1.Group("/user")
	userV1.POST("/login", r.Handler.User.Login)
	userV1.POST("/signup", r.Handler.User.SignUp)
	userV1.GET("/matching", r.Handler.User.FindUserMatching, r.Middleware.JWTMiddleware)
	userV1.GET("/match", r.Handler.User.FindUserMatch, r.Middleware.JWTMiddleware)

	swipeV1 := v1.Group("/swipe")
	swipeV1.POST("", r.Handler.Swipe.Swiping, r.Middleware.JWTMiddleware)

	packagesV1 := v1.Group("/packages")
	packagesV1.GET("", r.Handler.Packages.FindAllPackages)

	userPackageV1 := v1.Group("/user-package")
	userPackageV1.POST("/purchase", r.Handler.UserPackage.Purchase, r.Middleware.JWTMiddleware)
}
