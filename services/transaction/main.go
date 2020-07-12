package main

import (
	pbBalance "github.com/monewayTest/proto/balance/pb"
	pbTransaction "github.com/monewayTest/proto/transaction/pb"
	db "github.com/monewayTest/services/transaction/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

var balanceClient pbBalance.BalanceClient

var transactionDatabase db.TransactionDataBase

func main() {
	// Init the balance client
	con, client := getBalanceClient()
	balanceClient = client

	// Close the balance client at the end
	defer con.Close()

	// Init the transaction server
	lis, err := net.Listen("tcp", "localhost:4001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Close the transaction server at the end
	defer lis.Close()

	// Init the database
	tDatabase, err := db.NewDataBase()

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Close the database at the end
	defer tDatabase.Session.Close()

	transactionDatabase = *tDatabase
	g := grpc.NewServer()
	pbTransaction.RegisterTransactionServer(g, newServer())
	g.Serve(lis)
}