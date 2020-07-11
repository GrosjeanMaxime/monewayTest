#!/bin/bash

go build -o main_client main.go
go build -o account_service ./services/account/*.go
go build -o transaction_service ./services/transaction/*.go
go build -o balance_service ./services/balance/*.go
