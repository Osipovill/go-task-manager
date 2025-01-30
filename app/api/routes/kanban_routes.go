package routes

import (
	"app/exceptions"
	"app/models"
	"app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/tasks", GetTasks)
	r.GET("/tasks/:id", GetTaskByID)
	r.POST("/tasks", CreateTask)
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)
}

func GetTasks(c *gin.Context) {
	tasks, err := services.GetTasks()
	if err != nil {
		exceptions.HandleError(c, err, "Unable to fetch tasks")
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		exceptions.HandleError(c, err, "Invalid task ID")
		return
	}

	task, err := services.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		exceptions.HandleError(c, err, "Invalid input")
		return
	}

	createdTask, err := services.CreateTask(newTask.Title, newTask.Status, newTask.Priority)
	if err != nil {
		exceptions.HandleError(c, err, "Error creating task")
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		exceptions.HandleError(c, err, "Invalid task ID")
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		exceptions.HandleError(c, err, "Invalid input")
		return
	}

	task, err := services.UpdateTask(uint(id), updatedTask.Title, updatedTask.Status, updatedTask.Priority)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		exceptions.HandleError(c, err, "Invalid task ID")
		return
	}

	if err := services.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
