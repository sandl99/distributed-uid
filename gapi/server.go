package gapi

import (
	"context"
	"log"

	"github.com/sandl99/distributed-uid/snowflake"
)

type SnowFlakeServer struct {
	Snowflake snowflake.ISnowFlake
}

func NewSnowFlakeServer(snowflake snowflake.ISnowFlake) *SnowFlakeServer {
	return &SnowFlakeServer{Snowflake: snowflake}
}

func (server *SnowFlakeServer) Generate(ctx context.Context, req *UidRequest) (*UidResponse, error) {
	log.Printf("Recieve RPC request... \n")

	var uuid = server.Snowflake.Generate().Int64()
	var response = UidResponse{Uuid: uuid}
	return &response, nil
}

func (SnowFlakeServer) mustEmbedUnimplementedSnowFlakeServiceServer() {}
