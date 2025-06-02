package cmd

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

type Handl struct {
	DB *sql.DB
}

type Profile struct {
	Name string `json:"name"`
	id   int64  `json:"id"`
}

func (pr *Handl) Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("func profile: Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Println("Func profile0:failed to connetct", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println("Func profile1:failed to connetct", err)
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	}

	frs := cookie.Value
	log.Printf("cookie cookie:%v", frs)

	var name string
	var id int64
	err1 := pr.DB.QueryRow("SELECT name,id FROM person WHERE session_id = $1 LIMIT 1 ", frs).Scan(&name, &id)
	if err1 != nil {
		log.Println("Func profile2:failed to connetct", err1)
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	} else if err1 == sql.ErrNoRows {
		log.Println("Func profile3:failed to connetct", err1)
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	}

	slog.Info("Name", name)
	slog.Info("ID", id)

	//response := map[string]interface{}{
	//	"id":   userID,
	//	"name": NameFromBD,
	//}

	response := map[string]interface{}{
		"id":   id,
		"name": name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
