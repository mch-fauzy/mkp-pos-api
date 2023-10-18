package shared

import (
	"context"
	"net/http"

	"github.com/mkp-pos-cashier-api/shared/failure"
)

const (
	usernameKey = "username"
	roleKey     = "role"
	tokenKey    = "token"
)

// GetUsernameFromContext retrieves the username from the request context.
func GetUsernameFromContext(r *http.Request) (string, error) {
	username, ok := r.Context().Value(usernameKey).(string)
	if !ok {
		return "", failure.BadRequestFromString("Username not found in context")
	}
	return username, nil
}

func GetRoleFromContext(r *http.Request) (string, error) {
	role, ok := r.Context().Value(roleKey).(string)
	if !ok {
		return "", failure.BadRequestFromString("Role not found in context")
	}
	return role, nil
}

func GetTokenFromContext(r *http.Request) (string, error) {
	token, ok := r.Context().Value(tokenKey).(string)
	if !ok {
		return "", failure.BadRequestFromString("Token not found in context")
	}
	return token, nil
}

// WithUsername adds the username to the context.
func WithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, usernameKey, username)
}

func WithRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}
