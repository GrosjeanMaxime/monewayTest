#!/bin/bash

CREATE_MONEWAY_KEYSPACE="CREATE KEYSPACE IF NOT EXISTS moneway WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}"
CREATE_ACCOUNT_TABLE="CREATE TABLE IF NOT EXISTS moneway.accounts (
		id uuid PRIMARY KEY,
		name text,
		beneficiary text,
		iban text,
		bic text,
		create_at timestamp,
		updated_at timestamp,
		balance double )"
CREATE_TRANSACTION_TABLE="CREATE TABLE IF NOT EXISTS moneway.transactions (
		id text PRIMARY KEY,
		account_id text,
		description text,
		amount double,
		currency text,
		notes text,
		create_at timestamp,
		updated_at timestamp )"

until cqlsh -e "$CREATE_MONEWAY_KEYSPACE; $CREATE_ACCOUNT_TABLE; $CREATE_TRANSACTION_TABLE"; do
    echo "Unavailable"
    sleep 10
done &

exec /docker-entrypoint.py "$@"
