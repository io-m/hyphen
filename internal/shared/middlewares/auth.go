package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/io-m/hyphen/internal/shared/tokens"
	"github.com/io-m/hyphen/pkg/constants"
	"github.com/io-m/hyphen/pkg/helpers"
)

// MustAuthenticate is a middleware function that performs authentication.
func MustAuthenticate(protector tokens.IProtector) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Perform authentication logic here
			authorizationToken := r.Header.Get(constants.BEARER_TOKEN)
			fields := strings.Fields(authorizationToken)
			if len(fields) < 2 || !strings.EqualFold(fields[0], "Bearer") {
				helpers.ErrorResponse(w, errors.New("invalid authorization header format"), http.StatusUnauthorized)
				return
			}
			token := fields[1]
			claims, err := protector.VerifyToken(token)
			if err != nil {
				helpers.ErrorResponse(w, fmt.Errorf("token verification issue: %w", err), http.StatusUnauthorized)
				return
			}
			// Create a new request with the updated context
			r = r.WithContext(context.WithValue(r.Context(), constants.CLAIMS, claims))
			// Authentication successful, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
