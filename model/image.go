package model

import "time"

type Image struct {
	Name      string    `gorm:"primary_key" json:"name"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"updated_at"`
}
