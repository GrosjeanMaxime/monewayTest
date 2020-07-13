# monewayTest
The aim of the project is to simulate a banking transaction using different services (account, transaction, balance).
The different services communicate with each other using rGPC and protobuf technology.
The different data are stored in the ScyllaDb database.
The whole is managed with the Go language.

## Features

- Create an account
- Create / update transaction
- Get current balance

## Installation

Required : Dep, Go, Docker

Run the following command to install dependencies :
``` dep ensure ``` 

Install the ScyllaDb database :
``` sudo docker build -t scylladb . ``` 

Install the different services :
``` ./install_services.sh ```

Install the main client :
``` ./install_main_client.sh ```

## Usage

Run the ScyllaDb database :
``` sudo docker run -p 9042:9042 scylladb ```

Run the following commands in this order :
``` ./balance_service ./transaction_service ./account_service ``` 

Create an account :
``` ./main_client CREATE_ACCOUNT [NAME] [BENEFICIARY] ``` 

Create a transaction :
``` ./main_client CREATE_TRANSACTION [ACCOUNT_ID] [DESCRIPTION] [AMOUNT] [CURRENCY] [NOTES] ``` 

Update a transaction :
``` ./main_client UPDATE_TRANSACTION [TRANSACTION_ID] [DESCRIPTION] [CURRENCY] [NOTES] ``` 
 
Get the account balance :
``` ./main_client GET_BALANCE [ACCOUNT_ID] ``` 