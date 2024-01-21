package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"datink/internal/models/auth"
	"datink/pkg/util/httputil"
)

var bearer = "Bearer"

func (m Middlewares) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := len(bearer)
		tokenString := c.Request().Header.Get("Authorization")
		if len(tokenString) <= l+1 || !strings.EqualFold(tokenString[:l], bearer) {
			return m.Resource.Response.ErrorResponse(c, httputil.ErrorWithMessage(
				httputil.ErrUnauthorized,
				"unauthorized please login",
				fmt.Errorf("invalid auth header"),
			))
		}

		token, err := jwt.ParseWithClaims(tokenString[l+1:], &auth.UserJwtAuth{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.Resource.Config.JwtSigningKey), nil
		})

		if err != nil {
			return m.Resource.Response.ErrorResponse(c, httputil.ErrorWrap(httputil.ErrUnauthorized, err))
		}

		if claims, ok := token.Claims.(*auth.UserJwtAuth); ok && token.Valid {
			log.Println("claims", claims)
			c.Set("user", claims)
			return next(c)
		}

		return m.Resource.Response.ErrorResponse(c, httputil.ErrorWrap(httputil.ErrUnauthorized, fmt.Errorf("invalid auth header")))
	}
}
