package middleware

import (
	"fmt"
	"github.com/very-important-unmutable-organization/equipment/pkg/logger"
	"github.com/very-important-unmutable-organization/equipment/pkg/responses"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/render"
)

// Recoverer logs the panic and returns 500 status code & json body
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				logger.ErrorKV(r.Context(), "internal error",
					"panic", fmt.Sprintf("%+v", rvr),
					"stack", string(debug.Stack()),
				)

				render.Render(w, r, responses.ErrorInternal())
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
