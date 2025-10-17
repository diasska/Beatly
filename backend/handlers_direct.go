package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type DirectTask struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Date        string    `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func getTasksDirect(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlDB.Query("SELECT id, title, description, completed, date, created_at, updated_at FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []DirectTask
	for rows.Next() {
		var t DirectTask
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.Date, &t.CreatedAt, &t.UpdatedAt)
		tasks = append(tasks, t)
	}
	json.NewEncoder(w).Encode(tasks)
}

func getTaskDirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var t DirectTask
	err := sqlDB.QueryRow("SELECT id, title, description, completed, date, created_at, updated_at FROM tasks WHERE id = $1", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.Date, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(t)
}

func createTaskDirect(w http.ResponseWriter, r *http.Request) {
	var t DirectTask
	json.NewDecoder(r.Body).Decode(&t)
	res, err := sqlDB.Exec("INSERT INTO tasks (title, description, completed, date, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id",
		t.Title, t.Description, t.Completed, t.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	t.ID = int(id)
	json.NewEncoder(w).Encode(t)
}

func updateTaskDirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var t DirectTask
	json.NewDecoder(r.Body).Decode(&t)
	_, err := sqlDB.Exec("UPDATE tasks SET title=$1, description=$2, completed=$3, date=$4, updated_at=NOW() WHERE id=$5",
		t.Title, t.Description, t.Completed, t.Date, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ID = id
	json.NewEncoder(w).Encode(t)
}

func deleteTaskDirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	_, err := sqlDB.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Задача удалена"})
}
