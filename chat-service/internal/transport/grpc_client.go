package transport

import (
	"fmt"

	pb "github.com/SussyaPusya/swiftTalk/chat-service/client"
	"github.com/SussyaPusya/swiftTalk/chat-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	config *config.GRPC
	Client pb.Account_ServiceClient
	conn   *grpc.ClientConn
}

func NewClient(config *config.GRPC) *GRPCClient {

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}

	client := pb.NewAccount_ServiceClient(conn)

	return &GRPCClient{config: config, Client: client, conn: conn}
}

func (c *GRPCClient) Close() error {
	return c.conn.Close()
}
