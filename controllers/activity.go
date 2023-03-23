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

func GetAllActivityGroups(c *gin.Context){
	var activity []response.Result
	config.DB.Table("activities").Find(&activity)
	
	c.JSON(
		http.StatusOK,
        dtos.MediaDto{
            Status: "success",
			Message: "success",
			Data: activity,
        },

	)
}

func GetOneActivityGroups(c *gin.Context){
	var result response.Result
	fmt.Println(c.Param("id"))
	fmt.Println("ini", result)
	if err := config.DB.Table("activities").Where("id = ?", c.Param("id")).First(&result).Error; err != nil {
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

func CreateActivityGroups(c *gin.Context){
	var input form.Activity
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activity := models.Activity{
		Title: input.Title,
		Email: input.Email,
	}
	config.DB.Create(&activity)
	fmt.Println(activity.CreatedAt)

	// Create response
	response := response.Result{
		ID: activity.ID,
		Title: input.Title,
		Email: input.Email,
		CreatedAt: &activity.CreatedAt,
		UpdatedAt: &activity.UpdatedAt,
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

func UpdateActivityGroups(c *gin.Context){
	var activity models.Activity
	if err := config.DB.Table("activities").Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Activity
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&activity).Updates(input)

	// Create response
	curent_time := time.Now()
	response := response.Result{
		ID: activity.ID,
		Title: input.Title,
		Email: input.Email,
		CreatedAt: &activity.CreatedAt,
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

func DeleteActivityGroups(c *gin.Context){
	var activity models.Activity
	if err := config.DB.Table("activities").Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&activity)

	id := c.Param("id")
	message := fmt.Sprintf("Activity with ID %v Not Found", id)

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