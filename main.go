package main

import (
	"log/slog"
	"my-wallet-ntier-mongo/database"
	"my-wallet-ntier-mongo/response"
	"my-wallet-ntier-mongo/routes"
	"my-wallet-ntier-mongo/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load("config/env/.env"); err != nil {
		slog.Error("Error loading .env file")
	}

	database.ConnectDB()
}

func main() {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, response.ErrorMessage("Endpoint not found", 404))
	})

	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	v1 := e.Group("/v1")
	routes.UserRoute(v1)
	routes.TransactionRoute(v1)

	e.Logger.Fatal(e.Start(":1323"))
}
