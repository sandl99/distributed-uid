package main

import (
	"fmt"
	"net"

	"github.com/sandl99/distrbuted-uid/gapi"
	"github.com/sandl99/distrbuted-uid/snowflake"
	"google.golang.org/grpc"
)

func main() {
	snowflake.InfoLogger.Println("Starting the application...")
	port := snowflake.GetEnvInt64("PORT")

	constant, err := snowflake.InitConstant()
	if err != nil {
		snowflake.ErrorLogger.Fatal(err)
	}

	snowflaker, err := snowflake.NewSnowFlake(1, constant)
	if err != nil {
		snowflake.ErrorLogger.Fatal(err)
	}
	snowflake.InfoLogger.Printf("Starting the server on port %d ...\n", port)
	snowflakerServer := gapi.NewSnowFlakeServer(snowflaker)

	server := grpc.NewServer()
	gapi.RegisterSnowFlakeServiceServer(server, snowflakerServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		snowflake.ErrorLogger.Fatal("Can not start the server: ", err)
	}

	err = server.Serve(listener)
	if err != nil {
		snowflake.ErrorLogger.Fatal("Can not start the server: ", err)
	}
	snowflake.InfoLogger.Println("Stopping the application...")
}
