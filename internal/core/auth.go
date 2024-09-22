package core

import "context"

type (
	AuthService interface {
		Login(ctx context.Context, authConfig AuthConfig, user User) (*string, error)
		Signup(ctx context.Context, user User) error
	}

	AuthConfig struct {
		Secret   string
		TokenTTL int
	}
)
