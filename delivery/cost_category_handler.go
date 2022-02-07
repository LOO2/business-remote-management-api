package delivery

import (
	"strconv"

	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
	"github.com/gin-gonic/gin"
)

// represent the httphandler
type CostCategoryHandler struct {
	AUsecase models.CostCategoryUsecase
}

// will initialize the resources endpoint
func NewCostCategoryHandler(c *gin.Engine) {
	//handler := &CostCategoryHandler{}

	groupRoute := c.Group("/api")
	{
		revenueRoute := groupRoute.Group("cost_category")
		{
			revenueRoute.GET("/", ShowAllProviders)
		}
	}

}

func ShowAllCostCategories(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.CostCategory
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find revenue: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowCostCategory(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.CostCategory
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find cost category by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateCostCategory(c *gin.Context) {
	db := database.GetDatabase()

	var p models.CostCategory

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create cost category: " + err.Error(),
		})
		return
	}

	c.JSON(201, p)
}

func UpdateCostCategory(c *gin.Context) {
	db := database.GetDatabase()

	var p models.CostCategory

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create cost category: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteCostCategory(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.CostCategory{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete cost category: " + err.Error(),
		})
		return
	}

	c.Status(204)
}