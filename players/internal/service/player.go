package service

import (
	"context"
	"github.com/xcheng85/Go-EDA/players/internal/domain"
	"github.com/xcheng85/Go-EDA/players/internal/dto"
)

// Businiss logic regardless of api architecture
// DIP: ownership of interface
// define multiple custom type all at once
type (
	PlayerService interface {
		GetPlayer(ctx context.Context, get dto.GetPlayerRequestBody) (*domain.Player, error)
	}
	// only expose interface
	playerService struct {
	}
)

// do the interface checks
// if the AppImpl does not fulfill App interface, it will highlight
var _ PlayerService = (*playerService)(nil)

func NewPlayerService() PlayerService {
	return &playerService{}
}

func (p playerService) GetPlayer(ctx context.Context, get dto.GetPlayerRequestBody) (*domain.Player, error) {
	return &domain.Player{
		ID:   "0",
		Name: "dummy",
	}, nil
}
