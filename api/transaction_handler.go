package api

import (
	"log/slog"
	"my-wallet-ntier-mongo/response"
	"my-wallet-ntier-mongo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) GetTransactionTypes(c echo.Context) error {
	queryParams := response.TransactionTypeQuery{}
	if err := c.Bind(&queryParams); err != nil {
		slog.Error("GetTransactionTypes handler: cannot bind query param", "error", err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorMessage(err.Error(), http.StatusBadRequest))
	}

	if queryParams.Type != nil {
		if err := c.Validate(&queryParams); err != nil {
			slog.Error("Validation failed", "error", err.Error())
			return c.JSON(http.StatusBadRequest, response.ErrorMessage("Validation failed: type must be INCOME or OUTCOME only", http.StatusBadRequest))
		}
	}

	transactionTypes, total, err := h.transactionService.GetTransactionTypes(queryParams)
	if err != nil {
		slog.Error("GetTransactionTypes handler:", "error", err.Error())
		return c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
	}

	title := "Get transaction types"
	description := "Get transaction types is successfully."
	return c.JSON(http.StatusOK, response.SuccessMessage(response.DataObject{
		Title:       &title,
		Description: &description,
		Items:       &transactionTypes,
		Total:       &total,
	}))
}

func (h *TransactionHandler) GetTransactionsByUserId(c echo.Context) error {
	userId := c.Param("userId")
	if userId == "" {
		slog.Info("GetUserById handler: invalid userId")
		return c.JSON(http.StatusBadRequest, response.ErrorMessage("GetUserById handler: invalid userId", http.StatusBadRequest))
	}

	queryParams := response.TransactionQuery{}
	if err := c.Bind(&queryParams); err != nil {
		slog.Error("GetTransactionByUserId handler: cannot bind query param", "error", err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorMessage(err.Error(), http.StatusBadRequest))
	}
	if queryParams.Type != nil {
		if err := c.Validate(&queryParams); err != nil {
			slog.Error("Validation failed", "error", err.Error())
			return c.JSON(http.StatusBadRequest, response.ErrorMessage("Validation failed: type must be INCOME or OUTCOME only", http.StatusBadRequest))
		}
	}

	transactions, total, err := h.transactionService.GetTransactionsByUserId(userId, queryParams)
	if err != nil {
		slog.Error("GetTransactionByUserId handler:", "error", err.Error())
		return c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
	}

	title := "Get transaction by user id"
	description := "Get transaction by user id is successfully."
	return c.JSON(http.StatusOK, response.SuccessMessage(response.DataObject{
		Title:       &title,
		Description: &description,
		Items:       &transactions,
		Total:       &total,
	}))
}
