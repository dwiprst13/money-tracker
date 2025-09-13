package routes

import (
    "money-tracker/controllers"
    "money-tracker/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/transactions", controllers.CreateTransaction)
        auth.GET("/transactions", controllers.ListTransactions)
    }

    return r
}
