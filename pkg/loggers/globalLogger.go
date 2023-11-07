package loggers

import (
	"log"
	"os"
)

var GlobalLogger *log.Logger

func InitGlobalLogger() {
	file, err := os.OpenFile("/Users/romanovmaksim/GolandProjects/SaltAIdDishes/logs/globalLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	GlobalLogger = log.New(file, "", log.LstdFlags)
}
