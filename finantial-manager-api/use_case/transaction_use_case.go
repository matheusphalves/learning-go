package usecase

import (
	"finantial-manager-api/model"
	"finantial-manager-api/repository"
)

type TransactionUsecase struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCase(repository repository.TransactionRepository) TransactionUsecase {
	return TransactionUsecase{
		repository: repository,
	}
}

func (tu *TransactionUsecase) GetTransactions() ([]model.Transaction, error) {
	return tu.repository.GetTransactions()
}


func (tu *TransactionUsecase) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	
	id, err := tu.repository.CreateTransaction(transaction)

	if err != nil {
		return model.Transaction{}, err
	}

	transaction.ID = id

	return transaction, nil
}