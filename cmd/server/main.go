package main

import (
	"fmt"
	"net/http"
	"os"
	"simple-leave-tracker/internal/storage/db"
	"strconv"
)

func main() {
	_, err := db.New(DbDSN())
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}

func DbDSN() string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
	)

	return dsn
}
