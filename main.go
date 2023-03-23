package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"todolist-BE/controllers"
	"todolist-BE/config"
)

func main() {
	err := godotenv.Load(".env")
	if err!= nil {
        log.Fatal("Error loading environment")
    }
	port := os.Getenv("PORT")
	r := gin.Default()

	// Setup db
	config.SetupDatabaseConnection()
    
	// Activities
	r.GET("/activity-groups", controllers.GetAllActivityGroups)
	r.GET("/activity-groups/:id", controllers.GetOneActivityGroups)
	r.POST("/activity-groups", controllers.CreateActivityGroups)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivityGroups)
	r.DELETE("/activity-groups/:id", controllers.DeleteActivityGroups)

	// Todo
	// r.GET("/todo-items", controllers.GetAllTodoItems)
    // r.GET("/todo-items/:id", controllers.GetOneTodoItems)
    // r.POST("/todo-items", controllers.CreateTodoItems)
    // r.PATCH("/todo-items/:id", controllers.UpdateTodoItems)
    // r.DELETE("/todo-items/:id", controllers.DeleteTodoItems)

    log.Fatal(r.Run(":" + port))
}