package interceptor

import (
	"context"
	"strings"
	"user-service/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		//skip login and register
		if info.FullMethod == "/user.UserService/Login" ||
			info.FullMethod == "/user.UserService/CreateUser" {
			return handler(ctx, req)
		}

		//Extract metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		//get token
		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		tokenParts := strings.Split(authHeader[0], " ")
		if len(tokenParts) != 2 {
			return nil, status.Error(codes.Unauthenticated, "invalid token format")
		}

		tokenStr := tokenParts[1]

		//Validate JWT
		_, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "inalid token")
		}

		//continue to handler
		return handler(ctx, req)

	}
}
