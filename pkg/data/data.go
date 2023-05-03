package data

import "github.com/Jonas-Hoang/golang_reful_todolist/pkg/dto"

var Todo []dto.Todo

func init() {
	Todos = []dto.Todo{
		{ID: 1, Name: "Monday", Content: "Coding GoLang Suong suong"},
		{ID: 2, Name: "Tuesday", Content: "Coding GoLang Suong suong"},
		{ID: 1, Name: "Wenesday", Content: "Coding GoLang Suong suong"},
		{ID: 1, Name: "Thurday", Content: "Coding GoLang Suong suong"},
		{ID: 1, Name: "Friday", Content: "Coding GoLang Suong suong"},
	}
}
