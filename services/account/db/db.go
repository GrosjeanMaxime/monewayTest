package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"time"
)

type Account struct {
	Id          gocql.UUID
	Name        string
	Beneficiary string
	Iban        string
	Bic         string
	CreateAt    time.Time
	UpdatedAt   time.Time
	Balance     float64
}

type AccountDataBase struct {
	Session gocqlx.Session
}

func (a AccountDataBase) InsertAccount(account *Account) error {
	session := a.Session

	q :=  session.Query(qb.Insert("moneway.accounts").
		Columns("id", "name", "beneficiary", "iban", "bic", "create_at", "updated_at", "balance").ToCql())

	q.BindStruct(account)

	if err := q.ExecRelease(); err != nil {
		return err
	}
	return nil
}

func NewDataBase() (*AccountDataBase, error) {
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
		return &AccountDataBase{}, err
	}

	return &AccountDataBase{Session: session }, nil
}