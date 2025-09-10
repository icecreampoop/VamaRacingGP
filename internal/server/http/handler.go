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
