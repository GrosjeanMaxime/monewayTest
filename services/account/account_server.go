package main

import (
	"context"
	"github.com/GrosjeanMaxime/monewayTest/proto/account/pb"
	"github.com/GrosjeanMaxime/monewayTest/services/account/db"
	"github.com/gocql/gocql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type server struct{}

func newServer() *server {
	return &server{}
}

func (s *server) CreateAccount(ctx context.Context, account *pb.CreateRequestAccount) (*pb.ResponseAccount, error) {
	// Get random uuid for the account id
	accountId, _ := gocql.RandomUUID()

	newAccount := db.Account{
		Id:   accountId,
		Name:        account.Name,
		Beneficiary: account.Beneficiary,
		Iban:        "00000000000",
		Bic:         "0000",
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
		Balance:     0,
	}

	// Insert the account in the database
	err := accountDatabase.InsertAccount(&newAccount)

	if err != nil {
		return &pb.ResponseAccount{}, err
	}

	return  &pb.ResponseAccount{
		Id:			 newAccount.Id.String(),
		Name:        newAccount.Name,
		Beneficiary: newAccount.Beneficiary,
		Iban:        newAccount.Iban,
		Bic:         newAccount.Bic,
		CreateAt:  	 timestamppb.New(newAccount.CreateAt),
		UpdatedAt:   timestamppb.New(newAccount.CreateAt),
		Balance:     newAccount.Balance,
	}, nil
}