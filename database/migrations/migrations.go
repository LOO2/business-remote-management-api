package migrations

import (
	models "github.com/LOO2/business-remote-management-api/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Revenue{})
	db.AutoMigrate(models.CostCategory{})
	db.AutoMigrate(models.Provider{})
	db.AutoMigrate(models.Cost{})
	db.AutoMigrate(models.User{})
}
