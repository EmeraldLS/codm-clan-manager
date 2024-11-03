package middleware

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

const missingJWTErrorMessage = "Missing JWT token"
const invalidJWTErrorMessage = "Invalid JWT token"

func ValidateJWT(audience, domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerURL, err := url.Parse("https://" + domain + "/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid issuer URL"})
			c.Abort()
			return
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

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
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to set up JWT validator"})
			c.Abort()
			return
		}

		// Check for Bearer token
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": missingJWTErrorMessage})
			c.Abort()
			return
		}

		// Validate JWT
		_, err = jwtValidator.ValidateToken(c.Request.Context(), strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": missingJWTErrorMessage})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": invalidJWTErrorMessage})
			}
			c.Abort()
			return
		}

		// If everything is okay, proceed to the next handler
		c.Next()
	}
}
