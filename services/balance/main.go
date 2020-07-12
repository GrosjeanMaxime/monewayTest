package main

import (
	"github.com/monewayTest/proto/balance/pb"
	"github.com/monewayTest/services/balance/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

var accountDatabase db.AccountDataBase

func main() {
	// Init the balance server
	lis, err := net.Listen("tcp", "localhost:4002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Close the balance server at the end
	defer lis.Close()

	// Init the database
	aDatabase, err := db.NewDataBase()

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Close the database at the end
	defer aDatabase.Session.Close()

	accountDatabase = *aDatabase
	g := grpc.NewServer()

	pb.RegisterBalanceServer(g, newServer())
	g.Serve(lis)
}