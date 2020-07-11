package main

import (
	"context"
	"github.com/gocql/gocql"
	pbBalance "github.com/monewayTest/proto/balance/pb"
	pbTransaction "github.com/monewayTest/proto/transaction/pb"
	"github.com/monewayTest/services/transaction/db"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Server struct{}

func newServer() *Server {
	return &Server{}
}

func (s *Server) CreateTransaction(ctx context.Context, transaction *pbTransaction.CreateRequestTransaction) (*pbTransaction.ResponseTransaction, error) {
	transactionId, _ := gocql.RandomUUID()

	newTransaction := db.Transaction{
		Id:          transactionId.String(),
		AccountId:   transaction.GetAccountId(),
		Description: transaction.GetDescription(),
		Amount:      transaction.GetAmount(),
		Currency:    transaction.GetCurrency(),
		Notes:       transaction.GetNotes(),
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := balanceClient.UpdateBalance(context.Background(), &pbBalance.UpdateRequestBalance{AccountId: transaction.GetAccountId(), Amount: transaction.Amount})

	if err != nil {
		return &pbTransaction.ResponseTransaction{}, err
	}

	err = transactionDatabase.InsertTransaction(&newTransaction)

	return  &pbTransaction.ResponseTransaction{
		Id:          newTransaction.Id,
		AccountId:   newTransaction.AccountId,
		Description: newTransaction.Description,
		Amount:      newTransaction.Amount,
		Currency:    newTransaction.Currency,
		Notes:       newTransaction.Notes,
		CreateAt:    timestamppb.New(newTransaction.CreateAt),
		UpdatedAt:   timestamppb.New(newTransaction.UpdatedAt),
	}, nil
}

