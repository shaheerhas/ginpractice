package entity

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	ID         string  `json:"id" gorm:"primary_key"`
	Title      string  `json:"title"`
	ArtistName string  `json:"artist"`
	Price      float64 `json:"price"`
	// CreatedAt  time.Time `json:"created_at"`
	// UpdatedAt  time.Time `json:"updated_at"`
	// DeletedAt  time.Time `json:"deleted_at"`
}

/**/
