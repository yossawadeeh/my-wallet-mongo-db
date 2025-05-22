package routes

import (
	userHandler "my-wallet-ntier-mongo/api"
	"my-wallet-ntier-mongo/database"
	userRepo "my-wallet-ntier-mongo/repository"
	userService "my-wallet-ntier-mongo/service"

	"github.com/labstack/echo/v4"
)

func UserRoute(g *echo.Group) {
	newUserRepo := userRepo.NewUserRepository(database.MongoDB)
	newUserService := userService.NewUserService(newUserRepo)
	newUserHandler := userHandler.NewUserHandler(newUserService)

	userRoute := g.Group("/users")
	userRoute.GET("", newUserHandler.GetUsers)
	userRoute.GET("/:userId", newUserHandler.GetUserById)
}
