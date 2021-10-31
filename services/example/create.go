package example

import (
	"github.com/arifluthfi16/gomvcboilerplate/model"
)

func (s ExampleService) SayHi () string {
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