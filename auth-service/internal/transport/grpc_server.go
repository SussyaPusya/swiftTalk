package transport

import "google.golang.org/grpc"

type GRPCServer struct {
	handlers *GRPCHandlers
	server   *grpc.Server
}

func NewGRPCServer(handlers *GRPCHandlers) *GRPCServer {
	return &GRPCServer{handlers: handlers, server: grpc.NewServer()}
}

func (g *GRPCServer) Run() error {

	return nil
}

func (g *GRPCServer) ShutDown() {
	g.server.GracefulStop()
}
