package controller

import (
	"finantial-manager-api/model"
	"finantial-manager-api/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUsecase
}

func NewTransactionController(usecase usecase.TransactionUsecase) TransactionController {
	return TransactionController{
		transactionUseCase: usecase,
	}
}

func (transaction *TransactionController) GetTransactions(ctx *gin.Context) {

	transactions, err := transaction.transactionUseCase.GetTransactions()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (tr *TransactionController) CreateTransaction(ctx *gin.Context) {

	var transaction model.Transaction
	err := ctx.BindJSON(&transaction)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	newTransaction, err := tr.transactionUseCase.CreateTransaction(transaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newTransaction)
}
