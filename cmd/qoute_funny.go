package cmd

import (
	_ "context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Hbd struct {
	DB *sql.DB
}

func (d *Hbd) NewHbd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var p string

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("Func funnyrandom1:failed to connetct", err)
		return
	}

	var (
		qoute string
		auth  string
	)
	err := d.DB.QueryRow("SELECT qoute,auth FROM funny ORDER BY RANDOM() LIMIT 1 ").Scan(&qoute, &auth)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("Func funnyrandom2:failed to connetct", err)
		return
	} else if err == sql.ErrNoRows {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("Func funnyrandom3:failed to connetct", err)
		return
	}

	w.Header().Set("Content-Type", "aplication/json; charset=utf-8")

	resp := map[string]string{
		"qoutes": qoute,
		"auth":   auth,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("Func funnyrandom3:failed to connetct", err)
		return
	}

}
