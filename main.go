package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pbAccount "monewayTest/proto/account/pb"
	"os"
)

func createAccount(name string, beneficiary string) {
	conAccount, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conAccount.Close()

	client := pbAccount.NewAccountClient(conAccount)

	message, err := client.CreateAccount(context.Background(), &pbAccount.BasicAccount{
		Name:        name,
		Beneficiary: beneficiary,
	})

	if err != nil {
		log.Fatalf("%v, %v", client, err)
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
	default:
		fmt.Println("Bad request")
	}
}
