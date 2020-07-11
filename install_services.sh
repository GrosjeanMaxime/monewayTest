#!/bin/bash

go build -o main_client main.go
go build -o account_service ./services/account/*.go