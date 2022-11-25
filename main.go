package main

import (
	"hacktiv8-final-project-3/config"
	"hacktiv8-final-project-3/httpserver/controllers"
	"hacktiv8-final-project-3/httpserver/repositories"
	"hacktiv8-final-project-3/httpserver/routers"
	"hacktiv8-final-project-3/httpserver/services"
	"hacktiv8-final-project-3/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Environment Variables not found")
	}
	app := gin.Default()
	appRoute := app.Group("/api")
	db, _ := config.Connect()

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	routers.UserRouter(appRoute, userController, authService)

	app.Run(":3000")
}
