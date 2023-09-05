package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type InternalAuthInterceptor struct {
	apiKey         string
	ignoredMethods []string
}

func NewInternalAuthInterceptor(apiKey string, ignoredMethods []string) *InternalAuthInterceptor {
	return &InternalAuthInterceptor{
		apiKey:         apiKey,
		ignoredMethods: ignoredMethods,
	}
}

func (ai *InternalAuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		for _, m := range ai.ignoredMethods {
			if info.FullMethod == m {
				return handler(ctx, req)
			}
		}

		ctx, err := ai.authorize(ctx)
		if err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}

		return handler(ctx, req)
	}
}

func (ai *InternalAuthInterceptor) authorize(ctx context.Context) (context.Context, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(m["x-api-key"]) == 0 {
		return ctx, status.New(codes.Unauthenticated, "missing api key").Err()
	}

	if m["x-api-key"][0] != ai.apiKey {
		return ctx, status.New(codes.Unauthenticated, "unauthorized").Err()
	}

	return ctx, nil
}
