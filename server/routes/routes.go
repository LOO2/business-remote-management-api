package routes

import (
	"github.com/LOO2/business-remote-management-api/controllers"
	"github.com/LOO2/business-remote-management-api/delivery"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")
	{
		revenue := main.Group("revenue")
		{
			revenue.GET("/", delivery.ShowAllRevenues)
			revenue.GET("/:id", delivery.ShowRevenue)
			revenue.POST("/", delivery.CreateRevenue)
			revenue.PUT("/", delivery.UpdateRevenue)
			revenue.DELETE("/:id", delivery.DeleteRevenue)
		}

		cost_category := main.Group("cost_category")
		{
			cost_category.GET("/", controllers.ShowAllCostCategories)
			cost_category.GET("/:id", controllers.ShowCostCategory)
			cost_category.POST("/", controllers.CreateCostCategory)
			cost_category.PUT("/", controllers.UpdateCostCategory)
			cost_category.DELETE("/:id", controllers.DeleteCostCategory)
		}

		provider := main.Group("provider")
		{
			provider.GET("/", delivery.ShowAllProviders)
			provider.GET("/:id", delivery.ShowProviders)
			provider.POST("/", delivery.CreateProvider)
			provider.PUT("/", delivery.UpdateProvider)
			provider.DELETE("/:id", delivery.DeleteProvider)
		}
		cost := main.Group("cost")
		{
			cost.GET("/", controllers.ShowAllCost)
			cost.GET("/:id", controllers.ShowCost)
			cost.POST("/", controllers.CreateCost)
			cost.PUT("/", controllers.UpdateCost)
			cost.DELETE("/:id", controllers.DeleteCost)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.ShowAllUsers)
			user.GET("/:id", controllers.ShowUser)
			user.POST("/", controllers.CreateUser)
			user.PUT("/", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
		}
	}

	return router
}
