package data

import (
	"resful_api/pkg/dto"
)

var Todos []dto.Todo

func init() {
	Todos = []dto.Todo{
		{ID: 1, Name: "Monday", Content: "Coding GoLang Suong suong"},
		{ID: 2, Name: "Tuesday", Content: "Testing API"},
		{ID: 3, Name: "Wenesday", Content: "Post get delete ..."},
		{ID: 4, Name: "Thurday", Content: "Migrate database"},
		{ID: 5, Name: "Friday", Content: "Deploying GoLang"},
	}
}
