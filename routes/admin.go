package routes

import (
	"mows-game-center-time-mgt/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	//Game routes
	games := r.Group("/games")
	{
		games.POST("", handlers.AddNewGame)
		games.GET("", handlers.GetAllGames)
		games.GET("/:id", handlers.GetGamesByID)
		games.PUT("/update/:id", handlers.UpdateGameByID)
		games.GET("/title/:title", handlers.GetGameByName)
		games.DELETE("/:id", handlers.RemoveGame)
	}

	//Users routes
	users := r.Group("/users")
	{
		users.GET("", handlers.GetAllUsers)
		users.DELETE(":id",handlers.DeleteUser)
	}
	return r
}
