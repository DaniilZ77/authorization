package auth

import (
	"context"

	"github.com/DaniilZ77/authorization/internal/core"
)

func (s *server) getAuthConfigFromContext(ctx context.Context) (*core.AuthConfig, error) {
	secret, ok := ctx.Value(jwtSecretContextKey).(string)
	if !ok {
		return nil, core.ErrInvalidAuthConfig
	}

	tokenTTL, ok := ctx.Value(tokenTTLContextKey).(int)
	if !ok {
		return nil, core.ErrInvalidAuthConfig
	}

	return &core.AuthConfig{
		Secret:   secret,
		TokenTTL: tokenTTL,
	}, nil
}
