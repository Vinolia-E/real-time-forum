package main

import (
	"fmt"
	"net/http"

	"real_time_forum/backend/handlers"
	"real_time_forum/backend/sqlite"
)

func main() {
	err := sqlite.InitDatabase("real_time_database.db")
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("../frontend/style"))))

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
