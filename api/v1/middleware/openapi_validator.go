package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func OpenAPI(swagger *openapi3.T) func(next http.Handler) http.Handler {
	return nethttpmiddleware.OapiRequestValidatorWithOptions(swagger, &nethttpmiddleware.Options{
		Options: openapi3filter.Options{
			ExcludeRequestBody:          false,
			ExcludeRequestQueryParams:   false,
			ExcludeResponseBody:         false,
			ExcludeReadOnlyValidations:  false,
			ExcludeWriteOnlyValidations: false,
			IncludeResponseStatus:       true,
			MultiError:                  false,
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				return nil
			},
			SkipSettingDefaults: false,
		},
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			var err = struct {
				Error string `json:"error"`
			}{Error: message}
			json.NewEncoder(w).Encode(&err)
		},
		SilenceServersWarning: false,
	})
}
