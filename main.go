package main

import (
	"github.com/Pong-Kantapon/taskmanager/database"
	"github.com/Pong-Kantapon/taskmanager/models"
	"github.com/gin-gonic/gin"
)

func main() {

	//Connect Database
	database.ConnectDatabase()

	//Migrate task model
	database.DB.Statement.AutoMigrate(&models.Task{})

	//create Gin router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "APIs avaliable"})
	})

	//Localhost
	r.Run(":8080")

}
