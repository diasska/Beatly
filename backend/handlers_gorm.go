package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getTasksGORM(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func getTaskGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func createTaskGORM(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	db.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func updateTaskGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	db.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func deleteTaskGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	db.Delete(&task)
	json.NewEncoder(w).Encode(map[string]string{"message": "Задача удалена"})
}
