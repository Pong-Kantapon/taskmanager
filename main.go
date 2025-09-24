package main

import (
	"github.com/Pong-Kantapon/taskmanager/database"
	"github.com/Pong-Kantapon/taskmanager/models"
	"github.com/Pong-Kantapon/taskmanager/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//Connect Database
	database.ConnectDatabase()

	//Migrate task model
	database.DB.Statement.AutoMigrate(&models.Task{})

	//create Gin router
	r := gin.Default()

	routes.SetupRoutes(r)

	//Localhost
	r.Run(":8080")

}
