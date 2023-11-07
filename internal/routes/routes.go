package routes

import (
	"SaltAIdDishes/internal/handlers"
	"net/http"
)

func TESTRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", handlers.Home)

	//fileServer := http.FileServer(http.Dir("./web/static/"))
	//mux.Handle("/static", http.NotFoundHandler())
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

//======ТУТ ДОЛЖНА БЫТЬ ИЗОЛЯЦИЯ ДЛЯ ПАПКИ СТАТИК, НО Я ЕЕ НЕ ПОНЯЛ, СЛЕДОВАТЕЛЬНО НЕ ДОБАВИЛ
