package transport

import (
	"context"
	"errors"

	pb "github.com/SussyaPusya/swiftTalk/auth-service/api"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/pkg"
)

type GRPCHandlers struct {
	pb.UnimplementedAccount_ServiceServer
	j *pkg.JWTManager
}

func NewGRPCHandlers(j *pkg.JWTManager) *GRPCHandlers {
	return &GRPCHandlers{j: j}
}

func (s *GRPCHandlers) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{PingMessage: "PONG!!!!!"}, nil
}

func (s *GRPCHandlers) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	token := req.GetToken()
	if token == "" {
		return nil, errors.New("token is required")
	}

	claims, err := s.j.VerifyAccessToken(token)
	if err != nil {
		return nil, err
	}

	if claims.UserID == "" {
		return nil, nil
	}

	return &pb.ValidateTokenResponse{
		IsValid: true,
		UserId:  claims.UserID,
	}, nil
}
