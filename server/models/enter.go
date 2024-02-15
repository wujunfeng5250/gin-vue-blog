package models

import "time"

type MODEL struct {
	ID        uint      `json:"gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
