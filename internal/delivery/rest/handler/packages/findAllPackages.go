package packages

import (
	"github.com/labstack/echo/v4"

	_ "datink/internal/models/package/entity"
	_ "datink/pkg/util/httputil"
)

// FindAllPackages FindAllPackages
//
//	@ID			FindAllPackages
//	@Tags		Packages
//	@Produce	json
//	@Success	200					{object}	httputil.SuccessResponse{data=[]entity.Package}
//	@Failure	500					{object}	httputil.ErrorResponse{errors=[]string}
//	@Failure	400					{object}	httputil.ErrorResponse{errors=[]string}
//	@Router		/v1/user/packages	[get]
func (h Handler) FindAllPackages(ec echo.Context) error {
	ctx := ec.Request().Context()
	customers, err := h.Service.Package.FindAll(ctx)
	if err != nil {
		return h.Resource.Response.ErrorResponse(ec, err)
	}

	return h.Resource.Response.SuccessResponse(ec, "ok", customers)
}
