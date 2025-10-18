package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDB *sql.DB

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=ass2 port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных (GORM)")
	}

	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Note{})

	sqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Не удалось подключиться к базе данных (SQL)")
	}

	r := mux.NewRouter()

	r.Use(corsMiddleware)

	r.HandleFunc("/tasks", getTasksGORM).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTaskGORM).Methods("GET")
	r.HandleFunc("/tasks", createTaskGORM).Methods("POST")
	r.HandleFunc("/tasks/{id}", updateTaskGORM).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTaskGORM).Methods("DELETE")

	r.HandleFunc("/notes", getNotes).Methods("GET")
	r.HandleFunc("/notes/{id}", getNote).Methods("GET")
	r.HandleFunc("/notes", createNote).Methods("POST")
	r.HandleFunc("/notes/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")

	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		log.Printf("CORS: Обработка запроса %s от %s", r.Method, r.Header.Get("Origin"))

		next.ServeHTTP(w, r)
	})
}
