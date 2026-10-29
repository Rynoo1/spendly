package models

import "time"

type Category struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"type:varchar(50);not null;uniqueIndex"`
	Colour string `json:"colour" gorm:"type:varchar(7);not null;default: '#64748b'"`
	Icon   string `json:"icon" gorm:"type:varchar(50);not null;default: 'circle-dollar-sign'"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Transactions []Transaction `json:"-" gorm:"foreignKey:CategoryID"`
}
