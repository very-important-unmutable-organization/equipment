package middleware

import (
	"net/http"

	resp "github.com/very-important-unmutable-organization/equipment/pkg/responses"

	"github.com/go-chi/render"
)

func ApiKeyAuthentication(apiToken string, apiTokenHeader string) func(http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get(apiTokenHeader)
			if tokenHeader != apiToken {
				w.Header().Add("WWW-Authenticate", "Api Key Authentication")
				_ = render.Render(w, r, resp.ErrorUnauthorized())
				return
			}
			nextHandler.ServeHTTP(w, r)
		})
	}
}
