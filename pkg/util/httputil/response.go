package httputil

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type HttpResponse struct {
	logger *zap.Logger
}

var (
	ErrBadRequest     = buildError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, nil)
	ErrUnauthorized   = buildError(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, nil)
	ErrNotFound       = buildError(http.StatusText(http.StatusNotFound), http.StatusNotFound, nil)
	ErrInternalServer = buildError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, nil)
	ErrForbidden      = buildError(http.StatusText(http.StatusForbidden), http.StatusForbidden, nil)
)

func NewHttpResponse(logger *zap.Logger) *HttpResponse {
	return &HttpResponse{
		logger: logger,
	}
}

func (r *HttpResponse) SuccessResponse(ec echo.Context, message string, data interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusOK)
	}
	return ec.JSON(http.StatusOK, buildSuccessResponse(message, http.StatusOK, data))
}

func (r *HttpResponse) SuccessResponseWithCode(ec echo.Context, code int, message string, data interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusOK)
	}
	return ec.JSON(http.StatusOK, buildSuccessResponse(message, code, data))
}

func (r *HttpResponse) ErrorResponse(ec echo.Context, err error, request ...interface{}) error {
	httpErr, ok := err.(*HttpError)
	if !ok {
		httpErr = ErrorWrap(ErrInternalServer, err)
	}

	var req interface{}
	if len(request) > 0 {
		req = request[0]
	}

	r.logger.Error("", zap.Error(httpErr.err),
		zap.String("method", ec.Request().Method),
		zap.String("uri", ec.Request().RequestURI),
		zap.Any("request", req))

	return ec.JSON(httpErr.Code, buildErrorResponse(httpErr.Message, httpErr.Code, httpErr.err))
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func buildSuccessResponse(message string, code int, data interface{}) SuccessResponse {
	res := SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}

func buildErrorResponse(message string, code int, err error) ErrorResponse {
	errorMessage := err.Error()

	splitError := strings.Split(errorMessage, "\n")
	res := ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  splitError,
	}
	return res
}

type HttpError struct {
	ErrorResponse
	err error
}

func (e *HttpError) Error() string {
	return e.Message
}

func buildError(message string, status int, err error) *HttpError {
	if message == "" {
		message = http.StatusText(status)
	}
	if err == nil {
		err = errors.New("")
	}
	return &HttpError{
		ErrorResponse: buildErrorResponse(message, status, err),
		err:           err,
	}
}

func ErrorWrap(base *HttpError, err error) *HttpError {
	if base == nil {
		base = ErrInternalServer
	}
	return buildError(base.Message, base.Code, err)
}

func ErrorWithMessage(base *HttpError, message string, err error) *HttpError {
	if base == nil {
		base = ErrInternalServer
	}
	return buildError(message, base.Code, err)
}

func ErrorWithErrMessage(base *HttpError, err error) *HttpError {
	if base == nil {
		base = ErrInternalServer
	}
	message := ""
	if err != nil {
		message = err.Error()
	}
	return buildError(message, base.Code, err)
}
