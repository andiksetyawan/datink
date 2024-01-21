package user

import (
	"github.com/labstack/echo/v4"

	"datink/internal/models/user/dto"
	"datink/pkg/util/httputil"
)

// Login  Login user
//
//	@ID			Login
//	@Tags		User
//	@Produce	json
//	@Param		request			body		dto.LoginRequest	true	"request body"
//	@Success	200				{object}	httputil.SuccessResponse{data=dto.LoginResponse}
//	@Failure	500				{object}	httputil.ErrorResponse{errors=[]string}
//	@Failure	400				{object}	httputil.ErrorResponse{errors=[]string}
//	@Router		/v1/user/login	[post]
func (h Handler) Login(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := dto.LoginRequest{}
	if err := ec.Bind(&req); err != nil {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrBadRequest, err))
	}

	if err := ec.Validate(&req); err != nil {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrBadRequest, err), req)
	}

	result, err := h.Service.User.Login(ctx, req)

	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err, req)
	}

	return h.Resource.Response.SuccessResponse(ec, "success", result)
}
