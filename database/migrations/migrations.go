package migrations

import (
	models "github.com/LOO2/business-remote-management-api/Entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Revenue{})
	db.AutoMigrate(models.CostCategory{})
	db.AutoMigrate(models.Provider{})
	db.AutoMigrate(models.Cost{})
	db.AutoMigrate(models.User{})
}
