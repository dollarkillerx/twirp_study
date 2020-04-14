package module

import "github.com/jinzhu/gorm"

type Url struct {
	gorm.Model
	Key string `gorm:"key"`
	Url string `gorm:"url"`
}

