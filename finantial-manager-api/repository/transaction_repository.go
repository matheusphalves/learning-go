package repository

import (
	"database/sql"
	"finantial-manager-api/model"
	"fmt"
)

type TransactionRepository struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func (tr *TransactionRepository) GetTransactions() ([]model.Transaction, error) {
	query := "SELECT id, date, amount, type, description, month_year FROM public.transactions"

	rows, err := tr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Transaction{}, err
	}

	var transactionList []model.Transaction
	var transaction model.Transaction

	for rows.Next() {
		err = rows.Scan(
			&transaction.ID,
			&transaction.Date,
			&transaction.Amount,
			&transaction.Type,
			&transaction.Description,
			&transaction.MonthYear)

		if err != nil {
			fmt.Println(err)
			return []model.Transaction{}, err
		}

		transactionList = append(transactionList, transaction)
	}

	rows.Close()

	return transactionList, nil

}

func (tr *TransactionRepository) CreateTransaction(transaction model.Transaction) (int, error) {

	var id int

	query, err := tr.connection.Prepare(
		"INSERT INTO public.transactions" +
			"(date, amount, type, description, month_year) VALUES" +
			"($1, $2, $3, $4, $5) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	err = query.QueryRow(
		transaction.Date,
		transaction.Amount,
		transaction.Type,
		transaction.Description,
		transaction.MonthYear).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	query.Close()
	return id, nil
}

func (tr *TransactionRepository) GetTransactionById(id int) (*model.Transaction, error) {
	query, err := tr.connection.Prepare(
		"SELECT * FROM public.transactions WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var transaction model.Transaction

	err = query.QueryRow(id).Scan(
		&transaction.ID,
		&transaction.Date,
		&transaction.Amount,
		&transaction.Type,
		&transaction.Description,
		&transaction.MonthYear)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &transaction, nil
}

// func (tr *TransactionRepository) DeleteTransaction(transaction model.Transaction) (int, error) {
// 	query, err := tr.connection.Prepare(
// 		"DELETE FROM public.transactions where id = $1",
// 	)

// 	if err != nil {
// 		fmt.Println(err)
// 		return -1, err
// 	}
// }