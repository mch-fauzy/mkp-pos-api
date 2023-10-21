package middleware

import (
	"net/http"
	"strings"

	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/infras"
	"github.com/mkp-pos-cashier-api/shared"
	"github.com/mkp-pos-cashier-api/shared/failure"
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
			err := failure.Unauthorized("Missing token")
			response.WithError(w, err)
			return
		}

		// Parse and verify the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if token.Method.Alg() != jwt.SigningMethodHS256.Name {
				err := failure.InternalError(jwt.ErrSignatureInvalid)
				return nil, err
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				err := failure.InternalError(jwt.ErrInvalidKeyType)
				return nil, err
			}

			userRole, ok := claims[shared.RoleKey].(string)
			if !ok {
				err := failure.InternalError(jwt.ErrInvalidKeyType)
				return nil, err
			}

			userUsername, ok := claims[shared.UsernameKey].(string)
			if !ok {
				err := failure.InternalError(jwt.ErrInvalidKeyType)
				return nil, err
			}

			// Add role and username to context
			ctx := shared.WithRole(r.Context(), userRole)
			ctx = shared.WithUsername(ctx, userUsername)
			r = r.WithContext(ctx)

			return []byte(a.CFG.App.JWTAccessKey), nil
		})

		if err != nil {
			response.WithError(w, err)
			return
		}

		if !token.Valid {
			err = failure.Unauthorized("Invalid token")
			response.WithError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Authentication) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get role from context
		userRole, err := shared.GetRoleFromContext(r)
		if err != nil {
			response.WithError(w, err)
			return
		}

		if userRole != shared.AdminRole {
			err = failure.Forbidden("Forbidden")
			response.WithError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Authentication) IsCashier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get role from context
		userRole, err := shared.GetRoleFromContext(r)
		if err != nil {
			response.WithError(w, err)
			return
		}

		if userRole != shared.CashierRole {
			err = failure.Forbidden("Forbidden")
			response.WithError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
