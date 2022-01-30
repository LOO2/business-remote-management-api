package routes

import (
	"github.com/LOO2/business-remote-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		revenue := main.Group("revenue")
		{
			revenue.GET("/", controllers.ShowAllRevenues)
			revenue.GET("/:id", controllers.ShowRevenue)
			revenue.POST("/", controllers.CreateRevenue)
			revenue.PUT("/", controllers.UpdateRevenue)
			revenue.DELETE("/:id", controllers.DeleteRevenue)
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
			provider.GET("/", controllers.ShowAllProvider)
			provider.GET("/:id", controllers.ShowProviders)
			provider.POST("/", controllers.CreateProvider)
			provider.PUT("/", controllers.UpdateProvider)
			provider.DELETE("/:id", controllers.DeleteProvider)
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
