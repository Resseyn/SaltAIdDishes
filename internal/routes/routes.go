package routes

import (
	"SaltAIdDishes/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TESTRoute() (*http.ServeMux, *gin.Engine) {
	mux := http.NewServeMux()
	r := gin.Default()
	r.LoadHTMLGlob("web/html/*")
	r.Static("/static", "./web/static")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/", handlers.MainPage)
	r.POST("/", handlers.Post)
	r.POST("/generate", handlers.Generate)
	mux.HandleFunc("/postrecipe", handlers.Home)
	mux.HandleFunc("/post", handlers.YoutHome)
	mux.HandleFunc("/getlink", handlers.GetLink)

	//fileServer := http.FileServer(http.Dir("./web/static/"))
	//mux.Handle("/static", http.NotFoundHandler())
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux, r
}

//======ТУТ ДОЛЖНА БЫТЬ ИЗОЛЯЦИЯ ДЛЯ ПАПКИ СТАТИК, НО Я ЕЕ НЕ ПОНЯЛ, СЛЕДОВАТЕЛЬНО НЕ ДОБАВИЛ
