package model

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: Example{}},
	}
}
