package main

import (
	"context"
	pbBalance "github.com/GrosjeanMaxime/monewayTest/proto/balance/pb"
	pbTransaction "github.com/GrosjeanMaxime/monewayTest/proto/transaction/pb"
	"github.com/GrosjeanMaxime/monewayTest/services/transaction/db"
	"github.com/gocql/gocql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Server struct{}

func (s *Server) UpdateTransaction(ctx context.Context, transaction *pbTransaction.UpdateRequestTransaction) (*pbTransaction.ResponseTransaction, error) {
	// Parse the uuid for the transaction id
	uuid, _ := gocql.ParseUUID(transaction.Id)

	updateTransaction := db.Transaction{
		Id:          uuid,
		Description: transaction.GetDescription(),
		Currency:    transaction.GetCurrency(),
		Notes:       transaction.GetNotes(),
		UpdatedAt:   time.Now(),
	}

	// Update the transaction in the database
	err := transactionDatabase.UpdateTransaction(&updateTransaction)

	if err != nil {
		return &pbTransaction.ResponseTransaction{}, err
	}

	return  &pbTransaction.ResponseTransaction{
		Id:          updateTransaction.Id.String(),
		Description: updateTransaction.Description,
		Amount:      updateTransaction.Amount,
		Currency:    updateTransaction.Currency,
		Notes:       updateTransaction.Notes,
		UpdatedAt:   timestamppb.New(updateTransaction.UpdatedAt),
	}, nil
}

func (s *Server) CreateTransaction(ctx context.Context, transaction *pbTransaction.CreateRequestTransaction) (*pbTransaction.ResponseTransaction, error) {
	// Get random uuid for the transaction id
	transactionId, _ := gocql.RandomUUID()
	// Get parsed uuid for the account id
	accountId, _ := gocql.ParseUUID(transaction.AccountId)

	newTransaction := db.Transaction{
		Id:          transactionId,
		AccountId:   accountId,
		Description: transaction.GetDescription(),
		Amount:      transaction.GetAmount(),
		Currency:    transaction.GetCurrency(),
		Notes:       transaction.GetNotes(),
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Update the balance in the account database
	_, err := balanceClient.UpdateBalance(context.Background(), &pbBalance.UpdateRequestBalance{AccountId: transaction.GetAccountId(), Amount: transaction.Amount})

	if err != nil {
		return &pbTransaction.ResponseTransaction{}, err
	}

	// Insert the transaction in the database
	err = transactionDatabase.InsertTransaction(&newTransaction)

	return  &pbTransaction.ResponseTransaction{
		Id:          newTransaction.Id.String(),
		AccountId:   newTransaction.AccountId.String(),
		Description: newTransaction.Description,
		Amount:      newTransaction.Amount,
		Currency:    newTransaction.Currency,
		Notes:       newTransaction.Notes,
		CreateAt:    timestamppb.New(newTransaction.CreateAt),
		UpdatedAt:   timestamppb.New(newTransaction.UpdatedAt),
	}, nil
}

func newServer() *Server {
	return &Server{}
}