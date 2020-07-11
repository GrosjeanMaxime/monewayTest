package main

import (
	"google.golang.org/grpc"
	"log"
	"monewayTest/proto/account/pb"
	"monewayTest/services/account/db"
	"net"
)

var accountDatabase db.AccountDataBase

func main() {
	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	aDatabase, err := db.NewDataBase()

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	accountDatabase = *aDatabase
	g := grpc.NewServer()
	pb.RegisterAccountServer(g, newServer())
	g.Serve(lis)
}