package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getNotes(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlDB.Query("SELECT id, content FROM notes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, n)
	}
	json.NewEncoder(w).Encode(notes)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var n Note
	err := sqlDB.QueryRow("SELECT id, content FROM notes WHERE id=$1", id).Scan(&n.ID, &n.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Заметка не найдена", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(n)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var n Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := sqlDB.QueryRow("INSERT INTO notes (content) VALUES ($1) RETURNING id", n.Content).Scan(&n.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(n)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var n Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	n.ID = uint(id)

	_, err := sqlDB.Exec("UPDATE notes SET content=$1 WHERE id=$2", n.Content, n.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(n)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	_, err := sqlDB.Exec("DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
