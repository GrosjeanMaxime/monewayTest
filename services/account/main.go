package main

import (
	"github.com/GrosjeanMaxime/monewayTest/proto/account/pb"
	"github.com/GrosjeanMaxime/monewayTest/services/account/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

var accountDatabase db.AccountDataBase

func main() {
	// Init the account server
	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Close the account server at the end
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
	pb.RegisterAccountServer(g, newServer())
	g.Serve(lis)
}