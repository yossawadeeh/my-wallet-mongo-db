package routes

import (
	transactionHandler "my-wallet-ntier-mongo/api"
	"my-wallet-ntier-mongo/database"
	transactionRepo "my-wallet-ntier-mongo/repository"
	transactionService "my-wallet-ntier-mongo/service"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(g *echo.Group) {
	newTransactionRepo := transactionRepo.NewTransactionRepository(database.MongoDB)
	newTransactionService := transactionService.NewTransactionService(newTransactionRepo)
	newTransactionHandler := transactionHandler.NewTransactionHandler(newTransactionService)

	transactionRoute := g.Group("/transactions")
	transactionRoute.GET("/types", newTransactionHandler.GetTransactionTypes)
	transactionRoute.GET("/:userId", newTransactionHandler.GetTransactionsByUserId)
}
