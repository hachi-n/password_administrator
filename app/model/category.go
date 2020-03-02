package model

import (
	"fmt"
	"log"
)

func init() {
	createCategoryTable()
}

const categoryTableName = "categories"

type Category struct {
	ID   int
	Name string
}

func FindOrCreateByCategory(categoryName string) *Category {
	categories := SearchCategoriesByName(categoryName)
	if categories != nil {
		return categories[0]
	}
	category := CreateCategory(categoryName)
	return category
}

func SearchCategoriesByName(categoryName string) []*Category {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE name = ?`, categoryTableName)
	rows, err := DBConnection.Query(cmd, categoryName)
	defer rows.Close()
	if err != nil {
		log.Fatalf("category search query error. %v \n", err)
	}

	var categories []*Category
	for rows.Next() {
		category := new(Category)
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}
	return categories
}

func SearchCategoriesById(id int) *Category {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, categoryTableName)
	rows, err := DBConnection.Query(cmd, id)
	defer rows.Close()
	if err != nil {
		log.Fatalf("category search query error. %v \n", err)
	}

	var categories []*Category
	for rows.Next() {
		category := new(Category)
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}
	return categories[0]
}

func CreateCategory(categoryName string) *Category {
	cmd := fmt.Sprintf(`INSERT INTO %s ( name ) VALUES ( ? )`, categoryTableName)
	result, err := DBConnection.Exec(cmd, categoryName)
	if err != nil {
		log.Fatalf("category create query error. %v \n", err)
	}

	resultId, _ := result.LastInsertId()

	category := new(Category)
	category.ID = int(resultId)
	category.Name = categoryName

	return category
}
