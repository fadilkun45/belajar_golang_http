package repository

import (
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO customers(name) VALUES(?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)

	helper.PanicErr(err)

	id, err := result.LastInsertId()

	helper.PanicErr(err)

	category.Id = id

	return category
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)

	helper.PanicErr(err)

	return category

}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)

	helper.PanicErr(err)

}

func (c CategoryRepositoryImpl) FindByid(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)

	helper.PanicErr(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicErr(err)
		return category, nil
	} else {
		return category, errors.New("Category no found")
	}
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicErr(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicErr(err)
		categories = append(categories, category)
	}

	return categories
}
