package databaseModels

import (
	"SaltAIdDishes/internal/openAIdialog"
	"SaltAIdDishes/pkg/database"
	"SaltAIdDishes/pkg/loggers"
	"SaltAIdDishes/pkg/models"
	"database/sql"
	"github.com/lib/pq"
	"log"
	"time"
)

type DishesModel struct {
	DB *sql.DB
}

var Dishes DishesModel

func InitDishesModel() {
	Dishes.DB = database.GlobalDatabase
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
func (m *DishesModel) Translate() {
	rows, err := database.GlobalDatabase.Query("SELECT * FROM dishes WHERE russian_name IS NULL")
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, a, b, c, d, e string
		var empty interface{}
		var empt2 interface{}
		var updated bool
		updated = false

		err := rows.Scan(&id, &name, &a, &b, &c, &d, &e, &empt2, &empty)
		if err != nil {
			log.Fatal(err)
		}
		for !updated {
			translated, err := openAIdialog.Test("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", name)
			if err != nil {
				time.Sleep(50 * time.Second)
				continue
			} else {
				_, err := database.GlobalDatabase.Exec("UPDATE dishes SET russian_name = $1 WHERE id = $2", translated, id)
				if err != nil {
					loggers.ErrorLogger.Fatal(err)
				}
				updated = true
			}
		}
	}
}

//func (m *DishesModel) GetWithPatams(id int, params []string) (*models.Dish, error) {
//	var found *models.Dish
//
//	return found, nil
//}
