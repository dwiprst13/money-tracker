package routes

import (
	"github.com/gin-gonic/gin"
	"money-tracker/controllers"
	"money-tracker/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		
		api.GET("/transactions", controllers.GetTransactions)
		api.POST("/transactions", controllers.CreateTransaction)

		api.GET("/budgets", controllers.GetBudgets)
		api.POST("/budgets", controllers.CreateBudget)

		api.GET("/reports", controllers.GetReports)
	}
}
