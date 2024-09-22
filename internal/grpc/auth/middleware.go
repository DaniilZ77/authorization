package auth

import (
	"context"

	"google.golang.org/grpc"
)

func InjectJWTSecretInterceptor(jwtSecret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = context.WithValue(ctx, jwtSecretContextKey, jwtSecret)
		return handler(ctx, req)
	}
}

func InjectTokenTTLInterceptor(tokenTTL int) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = context.WithValue(ctx, tokenTTLContextKey, tokenTTL)
		return handler(ctx, req)
	}
}
