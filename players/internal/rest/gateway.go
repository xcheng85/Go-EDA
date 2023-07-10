package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/go-chi/chi/v5"
	"github.com/xcheng85/Go-EDA/players/playerspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// chi's router
// register the reverse proxy for grpc server
func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	// apiRoot matches api.annotations.yaml
	const apiRoot = "/api/players"

	gateway := runtime.NewServeMux()
	// two ways to register
	// 1. RegisterYourServiceHandlerFromEndpoint
	// 2. RegisterGreeterHandler
	err := playerspb.RegisterPlayersServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the GRPC gateway
	mux.Mount(apiRoot, gateway)

	return nil
}
