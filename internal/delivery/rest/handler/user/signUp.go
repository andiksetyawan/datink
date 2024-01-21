package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"datink/internal/models/user/dto"
	"datink/pkg/util/httputil"
)

// SignUp  SignUp user
//
//	@ID			SignUp
//	@Tags		User
//	@Produce	json
//	@Param		request			body		dto.SignUpRequest	true	"request body"
//	@Success	201				{object}	httputil.SuccessResponse{data=dto.SignUpResponse}
//	@Failure	500				{object}	httputil.ErrorResponse{errors=[]string}
//	@Failure	400				{object}	httputil.ErrorResponse{errors=[]string}
//	@Router		/v1/user/signup	[post]
func (h Handler) SignUp(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := dto.SignUpRequest{}
	if err := ec.Bind(&req); err != nil {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrBadRequest, err))
	}

	if err := ec.Validate(&req); err != nil {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrBadRequest, err), req)
	}

	result, err := h.Service.User.SignUp(ctx, req)

	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err, req)
	}

	return h.Resource.Response.SuccessResponseWithCode(ec, http.StatusCreated, "success", result)
}
