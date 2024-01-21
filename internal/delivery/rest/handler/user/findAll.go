package user

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) FindAll(ec echo.Context) error {
	ctx := ec.Request().Context()

	//_, ok := ctx.Value("user").(auth.UserJwtAuth)
	//if !ok {
	//	return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrUnauthorized, errors.New("unauthorized")))
	//}

	customers, err := h.Service.User.FindAll(ctx)
	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err)
	}

	return h.Resource.Response.SuccessResponse(ec, "ok", customers)
}
