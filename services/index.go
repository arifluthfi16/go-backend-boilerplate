package services

import (
	"github.com/arifluthfi16/gomvcboilerplate/services/example"
	"gorm.io/gorm"
)

var (
	ExampleService 	example.ExampleService
)

func InjectDBIntoServices (db *gorm.DB) {
	ExampleService.DB 	= db
}