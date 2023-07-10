package dto

import (
	//"github.com/xcheng85/Go-EDA/players/internal/domain"
	// "github.com/xcheng85/Go-EDA/players/playerspb"
)

type GetPlayerRequestBody struct {
	ID string `json:"id"`
}

// // for grpc api
// func CreateGetPlayerResponse(p *domain.Player) *playerspb.GetPlayerResponse {
// 	return &playerspb.GetPlayerResponse{
// 		Player: &playerspb.Player{
// 			Id:   p.ID,
// 			Name: p.Name,
// 		},
// 	}
// }
