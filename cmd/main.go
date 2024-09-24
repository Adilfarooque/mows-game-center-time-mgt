package main

import (
	"mows-game-center-time-mgt/cmd/handlers"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {
	r := gin.Default()
	r.GET("/games", handlers.GetAllGames)
	r.GET("/games/name/:name", handlers.GetGamesByName)
	r.POST("/games/create", handlers.AddNewGame)
	r.PUT("/games/update/:id", handlers.UpdateGame)
	r.DELETE("/games/delete/:id", handlers.RemoveGame)
	r.Run(port)
}
