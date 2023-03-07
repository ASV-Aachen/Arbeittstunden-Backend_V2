package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	First_name string
	Last_name  string
	Id         string
	Rolle      string
}

type Reduction struct {
	ID                   string
	Season               Season `gorm:"embedded"`
	Member               User   `gorm:"embedded"`
	Reduction_in_percent float32
	Note                 string
}

func (reduction *Reduction) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	reduction.ID = uuid.NewString()
	return
}
