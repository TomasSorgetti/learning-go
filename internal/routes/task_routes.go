package routes

import (
    "go-app/internal/handlers"
    "go-app/internal/repository"
    "go-app/internal/services"
    "database/sql"
    "github.com/gin-gonic/gin"
)

func SetupTaskRoutes(api *gin.RouterGroup, db *sql.DB) {
    taskRepo := repository.NewTaskRepository(db)
    taskService := services.NewTaskService(taskRepo)
    taskHandler := handlers.NewTaskHandler(taskService)

    tasks := api.Group("/tasks")
    {
        tasks.POST("", taskHandler.Create)
        tasks.GET("", taskHandler.GetAll)
        tasks.GET("/:id", taskHandler.GetByID)
        tasks.PUT("/:id", taskHandler.Update)
        tasks.DELETE("/:id", taskHandler.Delete)
    }
}