package configs

import (
	"UnderProject/internal/adapters/drivens/mysql"
	"UnderProject/internal/adapters/drivers/rest"
	"UnderProject/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//Repositories
	userRepo := mysql.NewUserRepositoryMySql()

	// Services
	userService := services.NewUserService(userRepo)

	// Handlers
	userHandler := rest.NewUserHandler(userService)

	// /api group
	api := router.Group("/api")
	{
		// /user group
		users := api.Group("/users")
		{
			users.POST("/", userHandler.UserCreate)
			users.GET("/:id", userHandler.UserGetById)
			users.GET("/", userHandler.UserGetAll)
			users.PUT("/:id", userHandler.UserUpdate)
			users.DELETE("/:id", userHandler.UserDelete)
		}
	}
}
