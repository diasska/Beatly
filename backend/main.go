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
	// Подключение к Postgres для GORM
	dsn := "host=localhost user=postgres password=kami333 dbname=DB port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных (GORM)")
	}

	// Автоматическая миграция модели Task
	db.AutoMigrate(&Task{})

	// Прямое подключение для SQL
	sqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Не удалось подключиться к базе данных (SQL)")
	}

	// Настройка роутера Gorilla Mux
	r := mux.NewRouter()

	// Middleware для CORS (применяется ко всем роутам)
	r.Use(corsMiddleware)

	// Endpoints через GORM
	r.HandleFunc("/tasks", getTasksGORM).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTaskGORM).Methods("GET")
	r.HandleFunc("/tasks", createTaskGORM).Methods("POST")
	r.HandleFunc("/tasks/{id}", updateTaskGORM).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTaskGORM).Methods("DELETE")

	// Endpoints через direct SQL
	r.HandleFunc("/direct/tasks", getTasksDirect).Methods("GET")
	r.HandleFunc("/direct/tasks/{id}", getTaskDirect).Methods("GET")
	r.HandleFunc("/direct/tasks", createTaskDirect).Methods("POST")
	r.HandleFunc("/direct/tasks/{id}", updateTaskDirect).Methods("PUT")
	r.HandleFunc("/direct/tasks/{id}", deleteTaskDirect).Methods("DELETE")

	// Добавьте обработчик для OPTIONS (на всякий случай, хотя middleware должен покрывать)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем origin для Vue (localhost:5173)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Если нужно куки/аутентификация

		// Для preflight (OPTIONS) возвращаем OK сразу
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Логируем для отладки
		log.Printf("CORS: Обработка запроса %s от %s", r.Method, r.Header.Get("Origin"))

		next.ServeHTTP(w, r)
	})
}
