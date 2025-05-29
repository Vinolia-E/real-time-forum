package main

import (
	"fmt"
	"real_time_forum/backend/sqlite"
)

func main() {
	err := sqlite.InitDatabase("real_time_database.db")
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
}
