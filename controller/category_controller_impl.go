package controller

import (
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/services"
	"belajar-golang-rest-api/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryServices
}

func NewCategoryController(categoryService services.CategoryServices) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (Controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := Controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webReponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webReponse)

}

func (Controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := Params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicErr(err)
	categoryUpdateRequest.Id = id

	categoryResponse := Controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webReponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webReponse)

}

func (Controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {

	categoryId := Params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicErr(err)

	Controller.CategoryService.Delete(request.Context(), id)

	webReponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webReponse)

}

func (Controller *CategoryControllerImpl) Find(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	categoryId := Params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicErr(err)

	categoryResponse := Controller.CategoryService.FindById(request.Context(), id)

	webReponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webReponse)

}

func (Controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	categoryResponses := Controller.CategoryService.FindAll(request.Context())

	webReponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webReponse)

}
