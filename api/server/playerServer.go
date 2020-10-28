package server

import (
	"context"

	pb "github.com/jessdwitch/tic-tac-micro/api/v1"
	"github.com/jessdwitch/tic-tac-micro/db"
)

type playerServer struct {
	pb.UnimplementedPlayerServiceServer
}

type gameServer struct {
	pb.UnimplementedGameServiceServer
}

func (s *playerServer) CreatePlayer(ctx context.Context, req *pb.CreatePlayerRequest) (*pb.Player, error) {
	return db.NewPlayer(req.Player.Name)
}
