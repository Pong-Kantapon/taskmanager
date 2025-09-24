package controllers

import (
	"net/http"

	"github.com/Pong-Kantapon/taskmanager/database"
	"github.com/Pong-Kantapon/taskmanager/models"
	"github.com/gin-gonic/gin"
)

// Get task
func GetTasks(c *gin.Context) {

	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// Create task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// Get task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if result := database.DB.First(&task, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Update task by ID
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if result := database.DB.First(&task, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&task).Updates(input)
	c.JSON(http.StatusOK, task)
}

// Delete task by ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if result := database.DB.First(&task, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
