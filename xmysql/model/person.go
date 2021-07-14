package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID       int       `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Age      int       `gorm:"column:age"`
	Birthday time.Time `gorm:"column:birthday"`
}

func (p *Person) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("fuck")
	return nil
}
