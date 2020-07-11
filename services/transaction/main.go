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
	con, client := getBalanceClient()
	balanceClient = client

	defer con.Close()

	lis, err := net.Listen("tcp", "localhost:4001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tDatabase, err := db.NewDataBase()

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	
	defer tDatabase.Session.Close()

	transactionDatabase = *tDatabase
	g := grpc.NewServer()
	pbTransaction.RegisterTransactionServer(g, newServer())
	g.Serve(lis)
}