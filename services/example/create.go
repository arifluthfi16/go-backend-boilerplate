package example

import (
	"fmt"
	"github.com/arifluthfi16/gomvcboilerplate/model"
)

func (s ExampleService) SayHi () string {
	fmt.Println("RUNNING AN EXAMPLE SERVICE")
	return "Successfully started the service"
}

func (s ExampleService) CreateOne (username string) (model.Example, error) {
	example := model.Example{
		Username:    username,
	}

	if err := s.DB.Create(&example).Error; err != nil {
		return example, err
	}
	return example, nil
}