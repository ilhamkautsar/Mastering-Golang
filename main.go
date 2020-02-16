package main

import (
	"fmt"
	"net/http"

	"github.com/ilham13/Mastering-Golang/basic"
)

func main() {
	// user := basic.Username{}
	// user.Cetak()

	users := basic.Users
	user := basic.User

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting web server at port 8000")
	http.ListenAndServe(":8000", nil)
}
