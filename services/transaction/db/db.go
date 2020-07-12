package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"time"
)

type Transaction struct {
	Id          gocql.UUID
	AccountId   gocql.UUID
	Description string
	Amount      float64
	Currency    string
	Notes       string
	CreateAt    time.Time
	UpdatedAt   time.Time
}

type TransactionDataBase struct {
	Session gocqlx.Session
}

func (a *TransactionDataBase) UpdateTransaction(transaction *Transaction) error {
	session := a.Session

	q:= session.Query(qb.Update("moneway.transactions").
		Set("description", "currency", "notes", "updated_at").
		Where(qb.Eq("id")).Existing().
		ToCql())

	q.BindStruct(transaction)

	if err := q.ExecRelease(); err != nil {
		return err
	}
	return nil
}

func (a *TransactionDataBase) InsertTransaction(transaction *Transaction) error {
	session := a.Session

	q :=session.Query(qb.Insert("moneway.transactions").
		Columns("id", "account_id", "description", "amount", "currency", "notes", "create_at", "updated_at").ToCql())

	q.BindStruct(transaction)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func NewDataBase() (*TransactionDataBase, error) {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min: time.Second,
		Max: 10 * time.Second,
		NumRetries: 5,
	}
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = gocql.Quorum
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	session, err := gocqlx.WrapSession((*cluster).CreateSession())

	if err != nil {
		return &TransactionDataBase{}, err
	}

	return &TransactionDataBase{Session: session }, nil
}