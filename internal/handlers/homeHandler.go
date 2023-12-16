package handlers

import (
	"SaltAIdDishes/internal/databaseModels"
	"SaltAIdDishes/internal/scrappers"
	"SaltAIdDishes/pkg/database"
	"SaltAIdDishes/pkg/loggers"
	"SaltAIdDishes/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.page.tmpl.html", gin.H{
		"title": "Main website",
	})
}
func Post(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title":  "Main website",
		"colors": fakeForm.Colors,
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://chat.openai.com")
	data := r.URL.Query().Get("data")
	dish := models.Dish{}
	dishRecipe := ""
	dataSplit := strings.Split(data, "    ")
	in := 1
	afterRecipe := false
	for _, st := range dataSplit {
		if st != "" {
			switch {
			case strings.Contains(st, "Название:"):
				dish.Name = strings.Trim(strings.Split(st, "Название:")[1], " ")
			case strings.Contains(st, "Краткое описание:"):
				dish.Description = strings.Trim(strings.Split(st, "Краткое описание:")[1], " ")
			case strings.Contains(st, "Описание:"):
				dish.Description = strings.Trim(strings.Split(st, "Описание:")[1], " ")
			case strings.Contains(st, "Ингредиенты:"):
				continue
			case strings.Contains(st, "Рецепт:"):
				dishRecipe = dishRecipe + st + "\n"
				afterRecipe = true
			default:
				if afterRecipe {
					st = fmt.Sprintf("  %v. ", in) + st
					in++
				} else if st[0:2] == "  " {
					if st[2:4] == "  " {
						st = st[0:2] + "  \u2022 " + st[4:]
					} else {
						st = "  \u2022 " + st[2:]
					}
				}
				dishRecipe = dishRecipe + st + "\n"
			}
		}
	}
	dish.Ingredients = strings.Split(dishRecipe, "Рецепт:")[0]
	dish.Recipe = strings.Split(dishRecipe, "Рецепт:")[1]
	tempparams := make([]string, 0, 10)

	tempparams = append(tempparams, "Итальянская")
	tempparams = append(tempparams, "Завтрак")
	tempparams = append(tempparams, "Закуска")
	tempparams = append(tempparams, "Дешевое")

	url := scrappers.Scrap(dish.Name)
	err := databaseModels.Dishes.Insert(dish.Name, dish.Description, dish.Ingredients, dish.Recipe, url, tempparams)
	if err != nil {
		//loggers.ErrorLogger.Println(err)
		fmt.Println("ИСПРАВЛЯЕТ")
		found, err := databaseModels.Dishes.Get(dish.Name)
		if err != nil {
			loggers.ErrorLogger.Println(err)
		}
		fmt.Println(found.Params)
		tempparams = append(tempparams, found.Params...)
		tempparams = removeDuplicates(tempparams)
		fmt.Println(tempparams)
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
func removeDuplicates(input []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, val := range input {
		if encountered[val] == false {
			encountered[val] = true
			result = append(result, val)
		}
	}

	return result
}
