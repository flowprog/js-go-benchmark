package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/gin-gonic/gin/binding"
	"stress/db"
	"stress/models"
)

func GetTasks(c *gin.Context) {
	tasks := []models.Task{
		{ID: "550e8400-e29b-41d4-a716-446655440000", Title: "Task 1", Completed: false},
		{ID: "550e8400-e29b-41d4-a716-446655440001", Title: "Task 2", Completed: true},
		{ID: "550e8400-e29b-41d4-a716-446655440002", Title: "Task 3", Completed: false},
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	var db = db.GetDB()

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&task)
	c.JSON(http.StatusOK, &task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(http.StatusOK, &task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&task)
}
