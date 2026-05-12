package transport

import (
	"fmt"
	"net"

	pb "github.com/SussyaPusya/swiftTalk/auth-service/api"
	"github.com/SussyaPusya/swiftTalk/auth-service/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	cfg      *config.GRPC
	handlers *GRPCHandlers
	server   *grpc.Server
	logger   *zap.Logger
}

func NewGRPCServer(handlers *GRPCHandlers, cfg *config.GRPC, logger *zap.Logger) *GRPCServer {
	return &GRPCServer{handlers: handlers, cfg: cfg, logger: logger}
}

func (g *GRPCServer) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", g.cfg.Port))
	if err != nil {
		g.logger.Error("failed to listen", zap.Error(err))
	}

	g.logger.Info("listening on", zap.String("addr", lis.Addr().String()))

	g.server = grpc.NewServer()

	pb.RegisterAccount_ServiceServer(g.server, g.handlers)

	if err := g.server.Serve(lis); err != nil {
		g.logger.Error("filed to serve", zap.Error(err))
	}
}

func (g *GRPCServer) ShutDown() {
	g.server.GracefulStop()
}
