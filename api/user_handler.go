package api

import (
	"log/slog"
	"my-wallet-ntier-mongo/constant"
	"my-wallet-ntier-mongo/interface/contract"
	"my-wallet-ntier-mongo/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService contract.UserService
}

func NewUserHandler(userService contract.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, total, err := h.userService.GetUsers()
	if err != nil {
		slog.Error("GetUsers handler:", "error", err.Error())
		return c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
	}

	title := "Get users"
	description := "Get users is successfully."
	return c.JSON(http.StatusOK, response.SuccessMessage(response.DataObject{
		Title:       &title,
		Description: &description,
		Items:       &users,
		Total:       &total,
	}))
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	userId := c.Param("userId")
	if userId == "" {
		slog.Info("GetUserById handler: invalid userId")
		return c.JSON(http.StatusBadRequest, response.ErrorMessage("GetUserById handler: invalid userId", http.StatusBadRequest))
	}

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		slog.Error("GetUserById handler:", "error", err.Error())

		switch err.Error() {
		case constant.INVALID_TYPE:
			return c.JSON(http.StatusBadRequest, response.ErrorMessage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		case constant.DOCUMENT_NOT_FOUND:
			return c.JSON(http.StatusNotFound, response.ErrorMessage(http.StatusText(http.StatusNotFound), http.StatusNotFound))
		default:
			return c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		}
	}

	title := "Get user"
	description := "Get user by id is successfully."
	return c.JSON(http.StatusOK, response.SuccessMessage(response.DataObject{
		Title:       &title,
		Description: &description,
		Item:        &user,
	}))
}
