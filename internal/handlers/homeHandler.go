package handlers

import (
	"SaltAIdDishes/internal/databaseModels"
	"SaltAIdDishes/internal/scrappers"
	"SaltAIdDishes/pkg/database"
	"SaltAIdDishes/pkg/loggers"
	"SaltAIdDishes/pkg/models"
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

//func MainPage(w http.ResponseWriter, r *http.Request)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://chat.openai.com")
	data := r.URL.Query().Get("data")
	dish := models.Dish{}
	dishRecipe := ""
	dataSplit := strings.Split(data, "    ")
	in := 1
	for _, st := range dataSplit {
		if st != "" {
			switch {
			case strings.Contains(st, "Название:"):
				dish.Name = strings.Trim(strings.Split(st, "Название:")[1], " ")
			case strings.Contains(st, "Краткое описание:"):
				dish.Description = strings.Trim(strings.Split(st, "Краткое описание:")[1], " ")
			case strings.Contains(st, "Описание:"):
				dish.Description = strings.Trim(strings.Split(st, "Описание:")[1], " ")
			default:
				if strings.Contains(st, "Ингредиенты:") {
					continue
				}
				if strings.Contains(st, "Рецепт:") {
					dishRecipe = dishRecipe + st + "\n"
				} else if !strings.Contains(st, "Рецепт:") {
					if st[0:2] == "  " {
						if st[2:4] == "  " {
							st = st[0:2] + "  \u2022 " + st[4:]
						} else {
							st = "  \u2022 " + st[2:]
						}
					} else if !strings.Contains(st, "Инструкции:") && st != "Ингредиенты:" {
						st = fmt.Sprintf("  %v. ", in) + st
						in++
					}
					dishRecipe = dishRecipe + st + "\n"
				}
			}
		}

	}
	dish.Ingredients = strings.Split(dishRecipe, "Рецепт:")[0]
	dish.Recipe = strings.Split(dishRecipe, "Рецепт:")[1]
	tempparams := make([]string, 0, 10)

	tempparams = append(tempparams, "Итальянская")
	tempparams = append(tempparams, "Завтрак")

	url := scrappers.Scrap(dish.Name)
	err := databaseModels.Dishes.Insert(dish.Name, dish.Description, dish.Ingredients, dish.Recipe, url, tempparams)
	if err != nil {
		//loggers.ErrorLogger.Println(err)
		fmt.Println("ИСПРАВЛЯЕТ")
		found, err := databaseModels.Dishes.Get(dish.Name)
		if err != nil {
			loggers.ErrorLogger.Println(err)
		}
		tempparams = append(tempparams, found.Params...)
		_, err = database.GlobalDatabase.Exec("UPDATE dishes SET params = $1 WHERE name = $2", pq.Array(tempparams), dish.Name)
		if err != nil {
			loggers.ErrorLogger.Println(err)
		}
	}
	found, err := databaseModels.Dishes.Get(dish.Name)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		panic(err)
	}
	fmt.Println(found.Name, found.Recipe)
}
