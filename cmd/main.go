package main

import (
	"fmt"
	"log"
	"mows-game-center-time-mgt/config"
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/routes"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {

	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading the config file")
	}
	fmt.Println(cfig)
	db, err := db.ConnectDatabase(cfig)
	if err != nil {
		log.Fatalf("Error connecting to the database:%v", err)
	}

	r := gin.Default()
	//userGroup := r.Group("/user")
	adminGroup := r.Group("/admin")
	//routes.UserRoutes(userGroup, db)
	routes.AdminRoutes(adminGroup, db)
	// r.GET("/games", handlers.GetAllGames)
	// r.GET("/games/name/:name", handlers.GetGamesByName)
	// r.POST("/games/create", handlers.AddNewGame)
	// r.PUT("/games/update/:id", handlers.UpdateGame)
	// r.DELETE("/games/delete/:id", handlers.RemoveGame)
	r.Run(port)
}
