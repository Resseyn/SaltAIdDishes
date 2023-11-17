package databaseModels

import (
	"SaltAIdDishes/internal/openAIdialog"
	"SaltAIdDishes/pkg/database"
	"SaltAIdDishes/pkg/loggers"
	"SaltAIdDishes/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"log"
	"strings"
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
	var params string
	err := m.DB.QueryRow("SELECT id, name, description, ingredients, recipe, url, params FROM dishes WHERE name = $1", name).Scan(&id, &name2, &description, &ingredients, &recipe, &url, &params)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return nil, err
	}
	params = strings.Trim(params, "{}")
	paramArray := strings.Split(params, ",")
	jsonString := fmt.Sprintf(`["%s"]`, strings.Join(paramArray, `","`))
	var result []string
	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return nil, err
	}
	found := &models.Dish{
		ID:          id,
		Name:        name,
		Description: description,
		Ingredients: ingredients,
		Recipe:      recipe,
		Url:         url,
		Params:      result,
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

func (m *DishesModel) GetWithParams(params []string) (*models.Dish, error) {
	found := models.Dish{}
	fmt.Println(params)
	var foundParams string
	if len(params) == 0 {
		err := database.GlobalDatabase.QueryRow("SELECT id, name, description, ingredients, recipe, url, params FROM dishes ORDER BY RANDOM() LIMIT 1").Scan(&found.ID, &found.Name, &found.Description, &found.Ingredients, &found.Recipe, &found.Url, &foundParams)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return &found, err
		}
		foundParams = strings.Trim(foundParams, "{}")
		paramArray := strings.Split(foundParams, ",")
		jsonString := fmt.Sprintf(`["%s"]`, strings.Join(paramArray, `","`))
		var result []string
		err = json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			fmt.Println("Ошибка декодирования JSON:", err)
			return &found, err
		}
		return &found, nil
	} else {
		//выбираем по all, но если нет совпадений, убираем последный элемент парамтров, также выводить флаг и предупреждать пользователя, если параметры подчистились
		err := database.GlobalDatabase.QueryRow("SELECT id, name, description, ingredients, recipe, url, params FROM dishes WHERE params @> $1 ORDER BY RANDOM() LIMIT 1", pq.Array(params)).Scan(&found.ID, &found.Name, &found.Description, &found.Ingredients, &found.Recipe, &found.Url, &foundParams)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return nil, err
		}
		foundParams = strings.Trim(foundParams, "{}")
		paramArray := strings.Split(foundParams, ",")
		jsonString := fmt.Sprintf(`["%s"]`, strings.Join(paramArray, `","`))
		var result []string
		err = json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			fmt.Println("Ошибка декодирования JSON:", err)
			return nil, err
		}
		return &found, nil
	}
}
