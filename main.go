package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New(".env File tidak ditemukan."))
		log.Fatal(".env File tidak ditemukan.")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT belum di setting."))
		log.Fatal("PORT belum di setting.")
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
