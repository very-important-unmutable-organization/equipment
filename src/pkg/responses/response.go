package responses

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Message string `json:"message"`
}

func OK() Response {
	return Response{Message: "OK"}
}

type ItemsResponse struct {
	Items interface{} `json:"items"`
}

type errorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	Error errorDetail `json:"error"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrorValidation(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Error: errorDetail{
			Code:    "validation_error",
			Message: err.Error(),
		},
	}
}

func ErrorTooManyRequests() render.Renderer {
	return &ErrorResponse{
		HTTPStatusCode: http.StatusTooManyRequests,
		Error: errorDetail{
			Code:    "too_many_requests",
			Message: http.StatusText(http.StatusTooManyRequests),
		},
	}
}

func ErrorNotFound(err error) render.Renderer {
	return &ErrorResponse{
		HTTPStatusCode: http.StatusNotFound,
		Error: errorDetail{
			Code:    "not_found",
			Message: http.StatusText(http.StatusNotFound),
		},
	}
}

func ErrorInternal() render.Renderer {
	return &ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Error: errorDetail{
			Code:    "internal_server_error",
			Message: http.StatusText(http.StatusInternalServerError),
		},
	}
}

func ErrorUnauthorized() render.Renderer {
	return &ErrorResponse{
		HTTPStatusCode: http.StatusUnauthorized,
		Error: errorDetail{
			Code:    "unauthorized",
			Message: http.StatusText(http.StatusUnauthorized),
		},
	}
}
