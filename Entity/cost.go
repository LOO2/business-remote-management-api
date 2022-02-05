package models

import (
	"time"

	"gorm.io/gorm"
)

type Cost struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	IDCostCategory uint           `json:"cost_category_id"`
	CostCategory   CostCategory   `gorm:"foreignKey:id;references:IDCostCategory" json:"category"`
	IDProvider     uint           `json:"provider_id"`
	CostProvider   Provider       `gorm:"foreignKey:id;references:IDProvider" json:"provider"`
	PurchaseDate   time.Time      `json:"purchase_date"`
	DueDate        time.Time      `json:"due_date"`
	Total          float32        `json:"total"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
