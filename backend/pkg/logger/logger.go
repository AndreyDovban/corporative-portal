package logger

import (
	"log"
	"os"
)

func INFO(data ...any) {

	file, err := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(data...)
	log.SetOutput(os.Stdout)
	log.Println(data...)

}

func ERROR(data ...any) {

	file, err := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("!!! Failed to open log file:", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(data...)
	log.SetOutput(os.Stdout)
	log.Println(data...)

}
