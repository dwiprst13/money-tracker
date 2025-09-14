package routes

import (
	"github.com/gin-gonic/gin"
	"money-tracker/controllers"
	"money-tracker/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.AuthMiddleware()) 

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/transactions", controllers.GetTransactions)
		api.POST("/transactions", controllers.CreateTransaction)
		api.PUT("/transactions/:id", controllers.UpdateTransaction)
		api.DELETE("/transactions/:id", controllers.DeleteTransaction)

		api.GET("/budgets", controllers.GetBudgets)
		api.POST("/budgets", controllers.CreateBudget)
		api.PUT("/budgets/:id", controllers.UpdateBudget)
		api.DELETE("/budgets/:id", controllers.DeleteBudget)

		api.GET("/reports", controllers.GetReports)
	}
}
