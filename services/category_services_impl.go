package services

import (
	exception "belajar-golang-rest-api/Exception"
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/model/domain"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/web"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServicesImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryServices {
	return &CategoryServicesImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServicesImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request)
	helper.PanicErr(err)

	tx, err := service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServicesImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request)
	helper.PanicErr(err)

	tx, err := service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServicesImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service *CategoryServicesImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServicesImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
