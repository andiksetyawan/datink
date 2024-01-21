package swipe

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"datink/internal/models/auth"
	"datink/internal/models/swipe/dto"
	"datink/pkg/util/httputil"
)

// Swiping  Swiping user
//
//	@ID			Swiping
//	@Tags		Swipe
//	@Produce	json
//	@Security	JWTAuth
//	@Param		request	body		dto.SwipingRequest	true	"request body"
//	@Success	200		{object}	httputil.SuccessResponse{data=dto.SwipingResponse}
//	@Failure	500		{object}	httputil.ErrorResponse{errors=[]string}
//	@Failure	400		{object}	httputil.ErrorResponse{errors=[]string}
//	@Router		/v1/swipe [post]
func (h Handler) Swiping(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := dto.SwipingRequest{}
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

	result, err := h.Service.Swipe.Swiping(ctx, user, req)

	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err, req)
	}

	return h.Resource.Response.SuccessResponse(ec, "success", result)
}
