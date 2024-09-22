package auth

type contextKey string

const (
	jwtSecretContextKey = contextKey("jwt_secret")
	tokenTTLContextKey  = contextKey("token_ttl")
)
