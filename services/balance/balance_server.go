package main

import (
	"context"
	"errors"
	"github.com/GrosjeanMaxime/monewayTest/proto/balance/pb"
)

type server struct{}

func (s *server) UpdateBalance(ctx context.Context, balanceRequest *pb.UpdateRequestBalance) (*pb.ResponseBalance, error) {
	// Get the balance in the account database
	balance, err := accountDatabase.GetBalance(balanceRequest.AccountId)

	if err != nil {
		return &pb.ResponseBalance{}, err
	}

	newBalance := balance + balanceRequest.Amount
	if newBalance < 0 {
		return &pb.ResponseBalance{}, errors.New("Not enough money in the account")
	}

	// Update the balance in the account database
	err = accountDatabase.UpdateBalance(balanceRequest.AccountId, newBalance)

	if err != nil {
		return &pb.ResponseBalance{}, err
	}

	return &pb.ResponseBalance{Amount: balance}, nil
}

func (s *server) GetBalance(ctx context.Context, account *pb.GetRequestBalance) (*pb.ResponseBalance, error) {
	// Get the balance in the account database
	balance, err := accountDatabase.GetBalance(account.AccountId)
	
	if err != nil {
		return &pb.ResponseBalance{}, err
	}
	
	return &pb.ResponseBalance{Amount: balance}, nil
}


func newServer() *server {
	return &server{}
}
