package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	// games := r.Group("/games")
	// {
	// 	games.POST("", handlers.AddNewGame)
	// 	//games.POST("/search")
	// }
	return r
}
