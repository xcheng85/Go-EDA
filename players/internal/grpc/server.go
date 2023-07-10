package grpc

import (
	"context"
	"github.com/xcheng85/Go-EDA/players/internal/dto"
	"github.com/xcheng85/Go-EDA/players/internal/service"
	"github.com/xcheng85/Go-EDA/players/playerspb"
	"google.golang.org/grpc"
)

// ref: https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/creating_main.go/
// server needs to implement interface: type PlayersServiceServer in api_grpc.pb.go
// controller has dependencies on business service
type server struct {
	playerService service.PlayerService
	playerspb.UnimplementedPlayersServiceServer
}

var _ playerspb.PlayersServiceServer = (*server)(nil)

// grpc.ServiceRegistrar: enables users to pass concrete types other than grpc.Server to the service
// Attach the Player service to the server
func RegisterServer(playerService service.PlayerService, registrar grpc.ServiceRegistrar) error {
	playerspb.RegisterPlayersServiceServer(registrar, server{playerService: playerService})
	return nil
}

func (s server) GetPlayer(ctx context.Context, request *playerspb.GetPlayerRequest) (*playerspb.GetPlayerResponse, error) {
	// convert from grpc request to dto
	player, err := s.playerService.GetPlayer(ctx, dto.GetPlayerRequestBody{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}
	// convert from domain/entity to grpc response
	return &playerspb.GetPlayerResponse{
		Player: &playerspb.Player{
			Id:   player.ID,
			Name: player.Name,
		},
	}, nil



	// return dto.CreateGetPlayerResponse(player), nil
}
