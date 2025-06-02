package main

import (
	"context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"

	"qoutes/cmd"
	"qoutes/iternal/database"
)

var Store = sessions.NewCookieStore([]byte("super-secret-key"))

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//front

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./Fronted/login.html")
			return
		}
		http.ServeFile(w, r, "./Fronted"+r.URL.Path)
	})

	db := database.Connect()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	h := &cmd.Handler{DB: db}
	p := &cmd.Handle{DB: db}
	e := &cmd.LoginCmd{DB: db}
	a := &cmd.Handl{DB: db}
	v := &cmd.Hbd{DB: db}

	http.HandleFunc("/login/api", e.Execute)
	http.HandleFunc("/register/api", p.Register)
	http.HandleFunc("/generate/api", h.ServeHTTP)
	http.HandleFunc("/profile/api", a.Profile)
	http.HandleFunc("/qoute/api", v.NewHbd)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
