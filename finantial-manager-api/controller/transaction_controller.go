package controller

import (
	"finantial-manager-api/model"
	"finantial-manager-api/use_case"
	"fmt"
	"net/http"
	"strconv"

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

func (tr *TransactionController) GetransactionById(ctx *gin.Context) {

	id := ctx.Param("id")

	if (id == "") {
		ctx.JSON(
			http.StatusBadRequest,
			model.Response{
				Message: "Id can´t be null",
			})
		return
	}

	transactionId, err := strconv.Atoi(id)

	if (err != nil) {
		ctx.JSON(
			http.StatusBadRequest,
			model.Response{
				Message: "Id should be a number",
			})
		return
	}

	transaction, err := tr.transactionUseCase.GetTransactionById(transactionId)

	if (transaction == nil) {
		ctx.JSON(
			http.StatusNotFound,
			model.Response{
				Message: fmt.Sprintf(
					"Transaction not found for id = %d", transactionId),
			})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}