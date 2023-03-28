package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Season struct {
	Year               int
	Obligatory_minutes int
}

type Project_item_hour struct {
	Id       string
	Member   User `gorm:"embedded"`
	Duration int
	Project_item Project_item `gorm:"embedded"`
}

func (member *Project_item_hour) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	member.Id = uuid.NewString()
	return
}

type Project_item struct {
	Id          string
	Project     Project `gorm:"embedded"`
	Season      Season  `gorm:"embedded"`
	Date        time.Time
	Title       string
	Description string
	Approved    bool
	Countable   bool
}

func (member *Project_item) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	member.Id = uuid.NewString()
	return
}

type Project struct {
	Id           string
	Name         string
	Description  string
	First_season Season `gorm:"embedded"`
	Last_season  Season `gorm:"embedded"`
}

func (member *Project) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	member.Id = uuid.NewString()
	return
}
