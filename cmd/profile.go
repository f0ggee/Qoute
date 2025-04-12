package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handl struct {
	DB *sql.DB
}

type Profile struct {
	Name string `json:"name"`
	ID   string `db:"id"`
}

func (pr *Handl) Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Printf("http.Error")
		return
	}
	var p *Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}

	cookies := cookie.Value

	log.Println("Func Rgister3:profile:", cookie.Value)

	var id string
	err = pr.DB.QueryRow("SELECT session_id FROM person where session_id = $1 ", cookie.Value).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}

	if id != cookies {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	log.Printf("Func Rgister3:profile: не найден")
	fmt.Printf("Пльзователь %s перенаправлен", id)
	http.Redirect(w, r, "/http://localhost:8080/register/api", http.StatusFound)

}
