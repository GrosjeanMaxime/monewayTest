package main

import (
	"github.com/monewayTest/proto/account/pb"
	"github.com/monewayTest/services/account/db"
	"google.golang.org/grpc"
	"log"
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

	defer aDatabase.Session.Close()

	accountDatabase = *aDatabase
	g := grpc.NewServer()
	pb.RegisterAccountServer(g, newServer())
	g.Serve(lis)
}