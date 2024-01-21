package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"datink/internal/models/auth"
	"datink/internal/models/user/dto"
	"datink/pkg/util/httputil"
)

// FindUserMatching FindUserMatching
//
//	@ID			FindUserMatching
//	@Tags		User
//	@Summary	Get List UserMatching
//	@Produce	json
//	@Security	JWTAuth
//	@Param		page				query		integer	false	"page"
//	@Param		limit				query		integer	false	"limit"
//	@Success	200					{object}	httputil.SuccessResponse{data=dto.FindUserMatchingResponse}
//	@Failure	500					{object}	httputil.ErrorResponse{errors=[]string}
//	@Failure	400					{object}	httputil.ErrorResponse{errors=[]string}
//	@Router		/v1/user/matching	[get]
func (h Handler) FindUserMatching(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := dto.FindUserMatchingRequest{}
	if err := ec.Bind(&req); err != nil {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrBadRequest, err))
	}

	if err := ec.Validate(&req); err != nil {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrBadRequest, err), req)
	}

	user, ok := ec.Get("user").(*auth.UserJwtAuth)
	if !ok {
		return h.Resource.Response.ErrorResponse(ec, httputil.ErrorWrap(httputil.ErrUnauthorized, errors.New("unauthorized")))
	}

	users, err := h.Service.User.FindUserMatching(ctx, user, req)
	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err)
	}

	return h.Resource.Response.SuccessResponse(ec, "ok", users)
}
