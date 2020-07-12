package main

import (
	"context"
	"fmt"
	pbAccount "github.com/GrosjeanMaxime/monewayTest/proto/account/pb"
	pbBalance "github.com/GrosjeanMaxime/monewayTest/proto/balance/pb"
	pbTransaction "github.com/GrosjeanMaxime/monewayTest/proto/transaction/pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

func getBalance(accountId string) {
	conBalance, err := grpc.Dial("localhost:4002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conBalance.Close()

	client := pbBalance.NewBalanceClient(conBalance)

	message, err := client.GetBalance(context.Background(), &pbBalance.GetRequestBalance{
		AccountId: accountId})

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(message)
}

func createTransaction(accountId string, description string, amount float64, currency string, notes string) {
	conAccount, err := grpc.Dial("localhost:4001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conAccount.Close()

	client := pbTransaction.NewTransactionClient(conAccount)

	message, err := client.CreateTransaction(context.Background(), &pbTransaction.CreateRequestTransaction{
		AccountId:   accountId,
		Description: description,
		Amount:      amount,
		Currency:    currency,
		Notes:       notes,
	})

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(message)
}

func updateTransaction(id string, description string, currency string, notes string) {
	conAccount, err := grpc.Dial("localhost:4001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conAccount.Close()

	client := pbTransaction.NewTransactionClient(conAccount)

	message, err := client.UpdateTransaction(context.Background(), &pbTransaction.UpdateRequestTransaction{
		Id:   id,
		Description: description,
		Currency:    currency,
		Notes:       notes,
	})

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(message)
}

func createAccount(name string, beneficiary string) {
	conAccount, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conAccount.Close()

	client := pbAccount.NewAccountClient(conAccount)

	message, err := client.CreateAccount(context.Background(), &pbAccount.CreateRequestAccount{
		Name:        name,
		Beneficiary: beneficiary,
	})

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(message)
}

func main() {
	query := os.Args[1:]

	switch {
	case len(query) == 0:
		fmt.Println("Bad request")
	case query[0] == "CREATE_ACCOUNT" && len(query) == 3:
		createAccount(query[1], query[2])
	case query[0] == "CREATE_TRANSACTION" && len(query) == 6:
		amount, _ := strconv.ParseFloat(query[3], 64)
		createTransaction(query[1], query[2], amount, query[4], query[5])
	case query[0] == "UPDATE_TRANSACTION" && len(query) == 5:
		updateTransaction(query[1], query[2], query[3], query[4])
	case query[0] == "GET_BALANCE" && len(query) == 2:
		getBalance(query[1])
	default:
		fmt.Println("Bad request")
	}
}
