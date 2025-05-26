package main

import (
	"finantial-manager-api/controller"
	"finantial-manager-api/repository"
	"finantial-manager-api/db"
	"finantial-manager-api/use_case"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	

	// Database
	dbConnection, err := db.ConnectDB()

	if (err != nil) {
		panic(err)
	}

	// Repository
	repository := repository.NewTransactionRepository(dbConnection)

	// Use Cases
	transactionUseCase := usecase.NewTransactionUseCase(repository)
	
	// Controllers
	transactionController := controller.NewTransactionController(transactionUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/transactions", transactionController.GetTransactions)
	server.POST("/transactions", transactionController.CreateTransaction)
	server.GET("transactions/:id", transactionController.GetransactionById)

	server.Run(":8000")
}