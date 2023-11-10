package main

import (
	"SaltAIdDishes/internal/database"
	"SaltAIdDishes/internal/routes"
	"SaltAIdDishes/pkg/loggers"
	"net/http"
)

func main() {
	loggers.InitErrorLogger()
	loggers.InitGlobalLogger()
	database.InitDatabase()
	database.InitDishesModel()
	go database.Translate()
	http.ListenAndServe(":80", routes.TESTRoute())
	loggers.GlobalLogger.Println("ZAPUSCAY")
}
