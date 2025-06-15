package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// Parse the HTML template
var tmpl = template.Must(template.ParseFiles("../frontend/index.html"))

var DB *sql.DB // assign this from main.go or init function

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl.Execute(w, nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	nickname := r.FormValue("nickname")
	age := r.FormValue("age")
	gender := r.FormValue("gender")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	cpassword := r.FormValue("cpassword")

	if password != cpassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	input := `INSERT INTO users (nick_name, age, gender, first_name, last_name, email, password_hash) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = DB.Exec(input, nickname, age, gender, fname, lname, email, string(hashedPassword))
	if err != nil {
		log.Println("DB Error:", err)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Registration successful",
	})
	// json.NewEncoder(w).Encode(map[string]string{
	// 	"status":  "ok",
	// 	"message": "Registration successful",
	// })
	// fmt.Fprintf(w, "Registration successful!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Login handler called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	var hashedPassword string

	err := DB.QueryRow("SELECT password_hash FROM users WHERE nick_name = ?", name).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// TODO: Validate against DB
	w.Write([]byte("Login successful!"))
}
