package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type CostCategory struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CostCategoryUsecase interface {
	GetByID(ctx context.Context, id int64) (CostCategory, error)
}

type CostCategoryRepository interface {
	GetByID(ctx context.Context, id int64) (CostCategory, error)
}
