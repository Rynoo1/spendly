package models

import "time"

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	AmountCents int64     `json:"amount_cents" gorm:"not null;check:amount_cents>0"`
	CategoryID  uint      `json:"category_id" gorm:"not null;index"`
	Date        time.Time `json:"date" gorm:"type:date;not null; index"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
