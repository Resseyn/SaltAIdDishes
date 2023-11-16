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
	//go databaseModels.Dishes.Translate()
	mux, r := routes.TESTRoute()
	go http.ListenAndServe(":80", mux)
	r.Run(":8080")
	loggers.GlobalLogger.Println("ZAPUSCAY")
}
