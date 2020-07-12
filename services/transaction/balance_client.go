package main

import (
	"github.com/GrosjeanMaxime/monewayTest/proto/balance/pb"
	"google.golang.org/grpc"
	"log"
)

func getBalanceClient() (*grpc.ClientConn, pb.BalanceClient)  {
	con, err := grpc.Dial("localhost:4002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	client := pb.NewBalanceClient(con)

	return con, client
}
