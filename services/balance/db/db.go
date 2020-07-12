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

func (a *AccountDataBase) GetBalance(accountId string) (float64, error) {
	session := a.Session

	// Set query to get balance according to the id
	q := session.Query(qb.Select("moneway.accounts").Where(qb.Eq("id")).ToCql())
	accountTmp := Account{}

	uuid, _ := gocql.ParseUUID(accountId)
	q.BindStruct(&Account{
		Id: uuid,
	})

	// Execute the query
	if err := q.GetRelease(&accountTmp); err != nil {
		return 0, err
	}
	return accountTmp.Balance, nil
}

func (a *AccountDataBase) UpdateBalance(accountId string, amount float64) error {
	session := a.Session

	// Set query to update balance according to the id
	q:= session.Query(qb.Update("moneway.accounts").
		Set("balance", "updated_at").
		Where(qb.Eq("id")).Existing().
		ToCql())

	uuid, _ := gocql.ParseUUID(accountId)
	q.BindStruct(&Account{
		Id: uuid,
		Balance: amount,
		UpdatedAt: time.Now(),
	})

	// Execute the query
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