package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/pkg/errors"
	"github.com/vought-esport-attendance/helpers"
)

// CustomClaims defines additional JWT claims if needed
type CustomClaims struct {
	Permissions []string `json:"permissions"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// ErrorMessage structure for JSON error responses
type ErrorMessage struct {
	Message string `json:"message"`
}

// ValidateJWT sets up the JWT middleware with custom error handling and claims validation
func ValidateJWT(audience, domain string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse the issuer URL from the domain
		issuerURL, err := url.Parse("https://" + domain + "/")
		if err != nil {
			log.Fatalf("Failed to parse the issuer URL: %v", err)
		}

		// Set up the JWT provider with caching
		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		// Initialize the JWT validator
		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{audience},
			validator.WithCustomClaims(func() validator.CustomClaims {
				return new(CustomClaims)
			}),
		)
		if err != nil {
			log.Fatalf("Failed to set up the JWT validator: %v", err)
		}

		// Check that the Authorization header is a Bearer token
		if authHeaderParts := strings.Fields(r.Header.Get("Authorization")); len(authHeaderParts) > 0 && strings.ToLower(authHeaderParts[0]) != "bearer" {
			helpers.WriteJSON(w, http.StatusUnauthorized, ErrorMessage{Message: "Bad credentials"})
			return
		}

		// Error handling for JWT validation
		errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Error validating JWT: %v", err)
			var status int
			var message string

			if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
				status = http.StatusUnauthorized
				message = "Requires authentication"
			} else if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
				status = http.StatusUnauthorized
				message = "Bad credentials"
			} else {
				status = http.StatusInternalServerError
				message = "Internal Server Error"
			}

			helpers.WriteJSON(w, status, ErrorMessage{Message: message})
		}

		// Set up the JWT middleware
		middleware := jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(errorHandler),
		)

		// Apply the middleware to the next handler
		middleware.CheckJWT(next).ServeHTTP(w, r)
	})
}
