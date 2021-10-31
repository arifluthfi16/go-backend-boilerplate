package example

import (
	"gorm.io/gorm"
)

type ExampleService struct {
	DB *gorm.DB
}

