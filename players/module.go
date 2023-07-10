package players

import (
	"context"
	"github.com/xcheng85/Go-EDA/internal/monolith"
	"github.com/xcheng85/Go-EDA/players/internal/grpc"
	"github.com/xcheng85/Go-EDA/players/internal/rest"
	"github.com/xcheng85/Go-EDA/players/internal/service"
	
)

type Module struct{}

// grpc global instance is passed in from mono
func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	playerService := service.NewPlayerService()
	if err := grpc.RegisterServer(playerService, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
