package routes

import (
	"mows-game-center-time-mgt/cmd/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup{
	 games := r.Group("/games")
	 {
		games.GET("",handlers.GetAllGames)
		games.POST("/search ")
	 }
	return r
}