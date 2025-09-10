package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Handler struct {
	DB *sql.DB
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}

type Motorcycle struct {
	ID         int     `json:"id"`
	Cost       float32 `json:"cost"`
	Power      float32 `json:"power"`
	Handling   float32 `json:"handling"`
	Durability float32 `json:"durability"`
}

func (h *Handler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
}

func (h *Handler) showAllFromUsersTable(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(SELECTALLFROMUSERTABLE)
	if err != nil {
		http.Error(w, "query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Balance)
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) showAllFromMotorcyclesTable(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(SELECTALLFROMMOTORCYCLESTABLE)
	if err != nil {
		http.Error(w, "query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var motors []Motorcycle

	for rows.Next() {
		var m Motorcycle
		rows.Scan(&m.ID, &m.Cost, &m.Power, &m.Handling, &m.Durability)
		motors = append(motors, m)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(motors)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
