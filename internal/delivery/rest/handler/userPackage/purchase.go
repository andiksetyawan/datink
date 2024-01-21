package userPackage

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"datink/internal/models/auth"
	"datink/internal/models/userPackage/dto"
	"datink/pkg/util/httputil"
)

// Purchase  Purchase User Package
//
//	@ID			Purchase
//	@Tags		UserPackage
//	@Produce	json
//	@Summary	Purchase User Premium Package
//	@Security	JWTAuth
//	@Param		request							body		dto.UserPackagePurchaseRequest	true	"request body"
//	@Success	201								{object}	httputil.SuccessResponse{data=dto.UserPackagePurchaseResponse}
//	@Failure	500								{object}	httputil.ErrorResponse{errors=[]string}
//	@Failure	400								{object}	httputil.ErrorResponse{errors=[]string}
//	@Router		/v1/user/user-package/purchase	[post]
func (h Handler) Purchase(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := dto.UserPackagePurchaseRequest{}
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

	result, err := h.Service.UserPackage.Purchase(ctx, user, req)
	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err, req)
	}

	return h.Resource.Response.SuccessResponseWithCode(ec, http.StatusCreated, "success", result)
}
