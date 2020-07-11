package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"time"
)

type Transaction struct {
	Id          string
	AccountId   string
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

func (a TransactionDataBase) InsertTransaction(account *Transaction) error {
	session := a.Session

	stmt, names := qb.Insert("moneway.transactions").
		Columns("id", "account_id", "description", "amount", "currency", "notes", "create_at", "updated_at").ToCql()
	insertTransaction := session.Query(stmt, names)

	insertTransaction.BindStruct(account)
	if err := insertTransaction.ExecRelease(); err != nil {
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