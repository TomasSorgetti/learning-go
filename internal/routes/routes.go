package routes

import (
    "database/sql"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
    api := r.Group("/api/v1")
    SetupTaskRoutes(api, db)
}