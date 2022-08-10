package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")

	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		logger.Log(errors.New("DB_USER belum di setting."))
		log.Fatal("DB_USER belum di setting.")
	}

	pass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		logger.Log(errors.New("DB_PASS belum di setting."))
		log.Fatal("DB_PASS belum di setting.")
	}

	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		logger.Log(errors.New("DB_HOST belum di setting."))
		log.Fatal("DB_HOST belum di setting.")
	}

	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, pass, host)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Koneksi Database Berhasil")
	}

	_, err = database.Exec(`CREATE DATABASE DB_GO_152308829101_69`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE DB_GO_152308829101_69`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE daftarTugas (
		    id INT AUTO_INCREMENT,
			petugas VARCHAR(250),
		    tugas TEXT NOT NULL,
			tenggatWaktu DATE,
		    status BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}
