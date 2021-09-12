package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Home ,%q", html.EscapeString(r.URL.Path))
	})
	//http://localhost:82/

	http.HandleFunc("/profile", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Profile, %q", html.EscapeString(r.URL.Path))
	})
	//http://localhost:82/profile

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Login, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})
	//http://localhost:82/login?username=pongchai&password=1234

	log.Fatal(http.ListenAndServe(":82", nil))
}
