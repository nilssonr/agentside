package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type Claims struct {
	UserID   string `json:"user_id"`
	TenantID string `json:"tenant_id"`
}

func (c Claims) Validate(ctx context.Context) error {
	if len(c.TenantID) == 0 {
		return fmt.Errorf("missing tenant_id in claims")
	}
	return nil
}

func Auth0(domain, audience string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		issuerURL, err := url.Parse(domain)
		if err != nil {
			log.Fatal(err)
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{audience},
			validator.WithCustomClaims(customClaims()),
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			log.Fatal(err)
		}

		middle := jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(errorHandler()),
		)

		return middle.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			validatedClaims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if !ok {
				http.Error(w, "failed to get validated claims", http.StatusInternalServerError)
				return
			}

			claims, ok := validatedClaims.CustomClaims.(*Claims)
			if !ok {
				http.Error(w, "failed to get custom claims", http.StatusInternalServerError)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), "tenant_id", claims.TenantID))
			r = r.WithContext(context.WithValue(r.Context(), "user_id", claims.UserID))

			next.ServeHTTP(w, r)
		}))
	}
}

func customClaims() func() validator.CustomClaims {
	return func() validator.CustomClaims {
		return &Claims{}
	}
}

func errorHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Failed to validate JWT"}`))
	}
}

func jwtHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
