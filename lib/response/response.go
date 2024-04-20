package response

import (
	"fmt"
	"math"
	"net/http"
)

const (
	StatusCodeSuccess             = "20000001"
	StatusCodeCreated             = "20100001"
	StatusCodeSuccessNotFoundData = "20000001"

	StatusCodeBadRequest          = "40000001"
	StatusCodeUnauthorized        = "40100001"
	StatusCodeForbidden           = "40300001"
	StatusCodeNotFound            = "40400001"
	StatusCodeConflict            = "40900001"
	StatusCodeUnprocessableEntity = "42200001"
	StatusCodeLocked              = "42300001"
	StatusCodeFailedDependency    = "42400001"
	StatusCodeTooEarly            = "42500001"
	StatusCodeUpgradeRequired     = "42600001"
	StatusCodeInternalServerError = "50000001"
)

const (
	ErrorBadRequest      = "error_binding_request"
	ErrorValidateRequest = "error_validation_request"
	ErrorGeneral         = "error_general"
	ErrorUnauthorized    = "error_unauthorized"
)

type ResponseStatus struct {
	Code    string      `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Errors  interface{} `json:"errors,omitempty"`
}

type ResponseMeta struct {
	CurrentPage int `json:"current_page" example:"1"`
	PerPage     int `json:"per_page" example:"10"`
	From        int `json:"from" example:"1"`
	To          int `json:"to" example:"10"`
	Total       int `json:"total" example:"100"`
	LastPage    int `json:"last_page" example:"10"`
}

type DefaultResponse struct {
	Status *ResponseStatus `json:"status,omitempty"`
	Meta   *ResponseMeta   `json:"meta,omitempty"`
	Data   interface{}     `json:"data,omitempty"`
}

type DefaultResponseNoData struct {
	Status *ResponseStatus `json:"status,omitempty"`
	Meta   *ResponseMeta   `json:"meta,omitempty"`
	Data   interface{}     `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status *ResponseStatus `json:"status,omitempty"`
	Meta   *ResponseMeta   `json:"meta,omitempty"`
	Data   interface{}     `json:"data,omitempty"`
}

// NewSuccessResponseCreated create new success payload
func NewSuccessResponseCreated(statusCode string, message string, data interface{}) DefaultResponse {
	return DefaultResponse{
		&ResponseStatus{
			Code:    statusCode,
			Message: message,
			Errors:  nil,
		},
		nil,
		data,
	}
}

// NewSuccessResponseStatusOK create new success payload
func NewSuccessResponseStatusOK(statusCode string, message string, data interface{}) DefaultResponse {
	return DefaultResponse{
		&ResponseStatus{
			Code:    statusCode,
			Message: message,
			Errors:  nil,
		},
		nil,
		data,
	}
}

// NewPaginationSuccessResponseStatusOK create new success payload
func NewPaginationSuccessResponseStatusOK(statusCode string, message string, data interface{}, meta ResponseMeta) DefaultResponse {
	return DefaultResponse{
		&ResponseStatus{
			Code:    statusCode,
			Message: message,
			Errors:  nil,
		},
		&meta,
		data,
	}
}

// NewSuccessResponseNoData create new success payload
func NewSuccessResponseNoData(statusCode string, message string) DefaultResponseNoData {
	return DefaultResponseNoData{
		&ResponseStatus{
			Code:    statusCode,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewBadRequestResponse bad request format response
func NewBadRequestResponse(err interface{}) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeBadRequest,
			Message: fmt.Sprintf("%v", err),
		},
		nil,
		nil,
	}
}

// NewInternalServerError default internalService server error with custom message
func NewInternalServerError(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeInternalServerError,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewForbiddenResponse default for Forbidden error response
func NewForbiddenResponse() ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeForbidden,
			Message: http.StatusText(http.StatusForbidden),
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewForbiddenResponseWithMessage default internalService server error with custom message
func NewForbiddenResponseWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeForbidden,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewNotFoundResponse default not found error response
func NewNotFoundResponse() ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeNotFound,
			Message: http.StatusText(http.StatusNotFound),
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewNotFoundResponseWithMessage default not found error response
func NewNotFoundResponseWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeNotFound,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewUnauthorized user unauthorized error response
func NewUnauthorized() ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewConflict default conflict entity error with custom message
func NewConflict(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeConflict,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewUnprocessableEntity default unprocessable entity error with custom message
func NewUnprocessableEntity(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeUnprocessableEntity,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewUnauthorizedWithMessage user unauthorized error response with custom message
func NewUnauthorizedWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeUnauthorized,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewLockededWithMessage locked error response with custom message
func NewLockededWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeLocked,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewFailedDependencyWithMessage failed dependency response with custom message
func NewFailedDependencyWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeFailedDependency,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewTooEarlyResponseWithMessage too early response with custom message
func NewTooEarlyResponseWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeTooEarly,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewUpgradeRequiredWithMessage upgrade required response with custom message
func NewUpgradeRequiredWithMessage(message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeUpgradeRequired,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewBadRequestWithMsg bad request format response
func NewBadRequestWithMsg(msg string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeBadRequest,
			Message: msg,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

// NewCustomCodeWithMessage default custom code error with custom message
func NewCustomCodeWithMessage(statusCode string, message string, code int) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    statusCode,
			Message: message,
			Errors:  nil,
		},
		nil,
		nil,
	}
}

func NewErrorValidate(message string, errors interface{}) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    StatusCodeUnprocessableEntity,
			Message: message,
			Errors:  errors,
		},
		nil,
		nil,
	}
}

func NewError(err error, code string, message string) ErrorResponse {
	return ErrorResponse{
		&ResponseStatus{
			Code:    code,
			Message: message,
			Errors:  err,
		},
		nil,
		nil,
	}
}

func (r *ResponseMeta) CountTotal(total, page, limit int) *ResponseMeta {
	r.Total = total
	from := page*limit - limit
	to := page * limit
	if to < 0 {
		to = 0
	}
	if from <= 0 {
		from = 1
	}
	r.To = to
	r.From = from

	lastpage := math.Ceil(float64(total) / float64(limit))
	if total < limit {
		lastpage = 1
	}
	r.LastPage = int(lastpage)
	r.CurrentPage = page
	r.PerPage = limit

	return r
}
