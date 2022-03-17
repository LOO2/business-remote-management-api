package routes

import (
	"github.com/LOO2/business-remote-management-api/delivery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	// same as
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

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
			cost_category.GET("/", delivery.ShowAllCostCategories)
			cost_category.GET("/:id", delivery.ShowCostCategory)
			cost_category.POST("/", delivery.CreateCostCategory)
			cost_category.PUT("/", delivery.UpdateCostCategory)
			cost_category.DELETE("/:id", delivery.DeleteCostCategory)
		}

		provider := main.Group("provider")
		{
			provider.GET("/", delivery.ShowAllProviders)
			provider.GET("/:id", delivery.ShowProvider)
			provider.POST("/", delivery.CreateProvider)
			provider.PUT("/", delivery.UpdateProvider)
			provider.DELETE("/:id", delivery.DeleteProvider)
		}
		cost := main.Group("cost")
		{
			cost.GET("/", delivery.ShowAllCosts)
			cost.GET("/:id", delivery.ShowCost)
			cost.POST("/", delivery.CreateCost)
			cost.PUT("/", delivery.UpdateCost)
			cost.DELETE("/:id", delivery.DeleteCost)
		}
		user := main.Group("user")
		{
			user.GET("/", delivery.ShowAllUsers)
			user.GET("/:id", delivery.ShowUser)
			user.POST("/", delivery.CreateUser)
			user.PUT("/", delivery.UpdateUser)
			user.DELETE("/:id", delivery.DeleteUser)
		}
	}

	return router
}
