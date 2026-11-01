package models

import (
	"encoding/json"
	"time"
)

type DateOnly struct {
	time.Time
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *DateOnly) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type Transaction struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Description string   `json:"description" gorm:"not null"`
	AmountCents int64    `json:"amount_cents" gorm:"not null;check:amount_cents>0"`
	CategoryID  uint     `json:"category_id" gorm:"not null;index"`
	Date        DateOnly `json:"date" gorm:"type:date;not null; index"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
