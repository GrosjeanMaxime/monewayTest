package main

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/monewayTest/proto/account/pb"
	"github.com/monewayTest/services/account/db"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type server struct{}

func newServer() *server {
	return &server{}
}

func (s *server) CreateAccount(ctx context.Context, account *pb.CreateRequestAccount) (*pb.ResponseAccount, error) {
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