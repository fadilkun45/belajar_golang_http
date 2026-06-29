package helper

import (
	"belajar-golang-rest-api/model/domain"
	"belajar-golang-rest-api/web"
)

func ToCategoryResponse(Category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   int(Category.Id),
		Name: Category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, Category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(Category))
	}
	return categoryResponses
}
