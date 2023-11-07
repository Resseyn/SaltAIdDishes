package database

import (
	"SaltAIdDishes/pkg/loggers"
	"SaltAIdDishes/pkg/models"
	"database/sql"
	"github.com/lib/pq"
)

type DishesModel struct {
	DB *sql.DB
}

var Dishes DishesModel

func InitDishesModel() {
	Dishes.DB = GlobalDatabase
}

func (m *DishesModel) Insert(name, description, ingredients, recipe, url string, params []string) error {
	dish := models.Dish{
		Name:        name,
		Description: description,
		Ingredients: ingredients,
		Recipe:      recipe,
		Url:         url,
		Params:      params,
	}
	_, err := m.DB.Exec("INSERT INTO dishes (name, description, ingredients, recipe, url, params) VALUES ($1, $2, $3, $4, $5, $6)",
		dish.Name, dish.Description, dish.Ingredients, dish.Recipe, dish.Url, pq.Array(dish.Params))
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}

	return nil
}

func (m *DishesModel) Get(name string) (*models.Dish, error) {
	var name2, description, ingredients, recipe, url string
	var id int
	err := m.DB.QueryRow("SELECT id, name, description, ingredients, recipe, url FROM dishes WHERE name = $1", name).Scan(&id, &name2, &description, &ingredients, &recipe, &url)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return nil, err
	}
	found := &models.Dish{
		ID:          id,
		Name:        name,
		Description: description,
		Ingredients: ingredients,
		Recipe:      recipe,
		Url:         url,
	}
	return found, nil
}

//func (m *DishesModel) GetWithPatams(id int, params []string) (*models.Dish, error) {
//	var found *models.Dish
//
//	return found, nil
//}
