package handlers

import (
	"SaltAIdDishes/pkg/database"
	"SaltAIdDishes/pkg/loggers"
	"fmt"
	"net/http"
	"strings"
)

var prevName string

func YoutHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://www.youtube.com")
	data := r.URL.Query().Get("data")
	fmt.Println(data)
	dataSplit := strings.Split(data, ".")
	link := "https://www.youtube.com" + dataSplit[1]
	name := dataSplit[0]
	if prevName != "" {
		fmt.Println("UPDATED", prevName, link)
		_, err := database.GlobalDatabase.Exec("UPDATE dishes SET link = $1 WHERE name = $2", link, prevName)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			fmt.Println(err)
		}
	}
	prevName = name
}
func GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://www.youtube.com")
	var nameOfNoLinker string
	err := database.GlobalDatabase.QueryRow("SELECT name FROM dishes WHERE link IS NULL").Scan(&nameOfNoLinker)
	if err != nil {
		fmt.Println(err)
		loggers.ErrorLogger.Println(err)
	}
	w.Write([]byte(nameOfNoLinker))
}
