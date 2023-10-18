package middleware

import (
	"net/http"
	"strings"

	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/infras"
	"github.com/mkp-pos-cashier-api/shared"
	"github.com/mkp-pos-cashier-api/transport/http/response"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	HeaderAuthorization = "Authorization"
)

type Authentication struct {
	DB  *infras.PostgreSQLConn
	CFG *configs.Config
}

func ProvideAuthentication(db *infras.PostgreSQLConn, cfg *configs.Config) *Authentication {
	return &Authentication{
		DB:  db,
		CFG: cfg,
	}
}

// Middleware for verifying JWT tokens
func (a *Authentication) VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the Authorization header
		authHeader := r.Header.Get(HeaderAuthorization)
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Check if the token is present
		if tokenString == "" {
			response.WithMessage(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Parse and verify the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if token.Method.Alg() != jwt.SigningMethodHS256.Name {
				return nil, jwt.ErrSignatureInvalid
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, jwt.ErrInvalidKeyType
			}

			userRole, ok := claims["role"].(string)
			if !ok {
				return nil, jwt.ErrInvalidKeyType
			}

			ctx := shared.WithRole(r.Context(), userRole)
			r = r.WithContext(ctx)

			return []byte(a.CFG.App.JWTAccessKey), nil
		})

		if err != nil || !token.Valid {
			response.WithMessage(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Authentication) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole, err := shared.GetRoleFromContext(r)
		if err != nil || userRole != shared.AdminRole {
			response.WithMessage(w, http.StatusForbidden, "Forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Authentication) IsCashier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole, err := shared.GetRoleFromContext(r)
		if err != nil || userRole != shared.CashierRole {
			response.WithMessage(w, http.StatusForbidden, "Forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}
