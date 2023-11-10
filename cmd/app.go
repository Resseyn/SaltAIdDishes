package main

import (
	"SaltAIdDishes/internal/databaseModels"
	"SaltAIdDishes/internal/routes"
	"SaltAIdDishes/pkg/database"
	"SaltAIdDishes/pkg/loggers"
	"net/http"
)

func main() {
	loggers.InitErrorLogger()
	loggers.InitGlobalLogger()
	database.InitDatabase()
	databaseModels.InitDishesModel()
	go databaseModels.Dishes.Translate()
	http.ListenAndServe(":80", routes.TESTRoute())
	loggers.GlobalLogger.Println("ZAPUSCAY")
}
