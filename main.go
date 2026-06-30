package main

import (
	exception "belajar-golang-rest-api/Exception"
	"belajar-golang-rest-api/app"
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/middleware"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryServiceImpl(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.Find)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicErr(err)
}
