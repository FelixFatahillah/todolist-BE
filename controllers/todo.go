package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"todolist-BE/config"
	"todolist-BE/response"
	"todolist-BE/dtos"
	"todolist-BE/form"
	"todolist-BE/models"
	"fmt"
	"time"
)

func GetAllTodoItems(c *gin.Context){
	var todo []response.TodoResult
	config.DB.Table("todos").Find(&todo)
	
	c.JSON(
		http.StatusOK,
        dtos.MediaDto{
            Status: "success",
			Message: "success",
			Data: todo,
        },

	)
}

func GetOneTodoItems(c *gin.Context){
	var result response.TodoResult
	if err := config.DB.Table("todos").Where("id = ?", c.Param("id")).First(&result).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(
		http.StatusOK,
        dtos.MediaDto{
            Status: "success",
			Message: "success",
			Data: result,
        },

	)
}

func CreateTodoItems(c *gin.Context){
	var input form.Todo
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("isactive  = ", input.IsActive)
	todo := models.Todos{
		ActivityGroupId: input.ActivityGroupId,
		Title: input.Title,
		IsActive: input.IsActive,
		Priority: input.Priority,
	}
	config.DB.Create(&todo)

	// Create response
	response := response.TodoResult{
		ID: todo.ID,
		ActivityGroupId: input.ActivityGroupId,
		Title: input.Title,
		IsActive: input.IsActive,
		Priority: input.Priority,
		CreatedAt: &todo.CreatedAt,
		UpdatedAt: &todo.UpdatedAt,
	}

	c.JSON(
		http.StatusOK,
        dtos.MediaDto{
            Status: "success",
			Message: "success",
			Data: response,
        },

	)
}

func UpdateTodoItems(c *gin.Context){
	var todo models.Todos
	if err := config.DB.Table("todos").Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.TodoUpdate
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&todo).Updates(input)

	// Create response
	curent_time := time.Now()
	response := response.TodoResult{
		ID: todo.ID,
		ActivityGroupId: input.ActivityGroupId,
		Title: input.Title,
		IsActive: input.IsActive,
		Priority: input.Priority,
		CreatedAt: &todo.CreatedAt,
		UpdatedAt: &curent_time,
	}

	c.JSON(
		http.StatusOK,
        dtos.MediaDto{
            Status: "success",
			Message: "success",
			Data: response,
        },

	)
}

func DeleteTodoItems(c *gin.Context){
	var todo models.Todos
	if err := config.DB.Table("activities").Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&todo)

	id := c.Param("id")
	message := fmt.Sprintf("Todo with ID %v Not Found", id)

	type Response struct {
		Status string `json:"status"`
		Message string `json:"message"`
	}

	c.JSON(
		http.StatusOK,
        Response{
            Status: "Not Found",
			Message: message,
        },

	)
}