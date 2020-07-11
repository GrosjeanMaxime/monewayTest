package main

import (
	"github.com/gocql/gocql"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"monewayTest/proto/account/pb"
	"monewayTest/services/account/db"
	"time"
)

type server struct{}

func newServer() *server {
	return &server{}
}

func (s *server) CreateAccount(ctx context.Context, account *pb.BasicAccount) (*pb.CompleteAccount, error) {
	accountId, _ := gocql.RandomUUID()

	newAccount := db.Account{
		Id:   accountId.String(),
		Name:        account.Name,
		Beneficiary: account.Beneficiary,
		Iban:        "00000000000",
		Bic:         "0000",
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
		Balance:     0,
	}
	err := accountDatabase.InsertAccount(&newAccount)

	return  &pb.CompleteAccount{
		Id:			 newAccount.Id,
		Name:        newAccount.Name,
		Beneficiary: newAccount.Beneficiary,
		Iban:        newAccount.Iban,
		Bic:         newAccount.Bic,
		CreateAt:  	 timestamppb.New(newAccount.CreateAt),
		UpdatedAt:   timestamppb.New(newAccount.CreateAt),
		Balance:     newAccount.Balance,
	}, err
}