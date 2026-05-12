package transport

import (
	"context"

	pb "github.com/SussyaPusya/swiftTalk/auth-service/api"
)

type GRPCHandlers struct {
	pb.UnimplementedAccount_ServiceServer
}

func NewGRPCHandlers() *GRPCHandlers {
	return &GRPCHandlers{}
}

func (s *GRPCHandlers) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{PingMessage: "PONG!!!!!"}, nil
}
