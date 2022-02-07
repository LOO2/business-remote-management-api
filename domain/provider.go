package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Cnpj      int            `json:"cnpj"`
	Telephone int            `json:"telephone"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// RevenueUsecase represent the revenue's usecases
type ProviderUsecase interface {
	GetByID(ctx context.Context, id int64) (Revenue, error)
}

// RevenueRepository represent the revenue's repository contract
type ProviderRepository interface {
	GetByID(ctx context.Context, id int64) (Revenue, error)
}
