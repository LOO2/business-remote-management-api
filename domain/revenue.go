package models

import (
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
