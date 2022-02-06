package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Revenue struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Voucher   float32        `json:"voucher"`
	Debit     float32        `json:"debit"`
	Credit    float32        `json:"credit"`
	Delivery  float32        `json:"delivery"`
	Total     float32        `json:"total"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

// RevenueUsecase represent the revenue's usecases
type RevenueUsecase interface {
	GetByID(ctx context.Context, id int64) (Revenue, error)
}

// RevenueRepository represent the revenue's repository contract
type RevenueleRepository interface {
	GetByID(ctx context.Context, id int64) (Revenue, error)
}
