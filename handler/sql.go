package handler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
}

func InitDB() {
	db, err := sql.Open("sqlite3", "brew.db")
	if err != nil {
			fmt.Println("Error opening SQLite database:", err)
			return
	}
	defer db.Close()

		// Create the "users" table if it does not exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		fmt.Println("Error creating users table:", err)
		return
	}

	fmt.Println("User table created or already exists.")
}

func UsersCreate(username, password string) error{
	db, err := sql.Open("sqlite3", "brew.db")
	if err != nil {
			fmt.Println("Error opening SQLite database:", err)
			return err
	}
  defer db.Close()

	userID := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
			fmt.Println("Error hashing password", err)
			return err
	}

	newUser := User{
		ID:        userID,
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	_, err = db.Exec(`
		INSERT INTO users (id, username, password, created_at)
		VALUES (?, ?, ?, ?)
	`, newUser.ID, newUser.Username, newUser.Password, newUser.CreatedAt)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	fmt.Println("User created successfully.")
	return nil
}

func UsersLogin(username, password string) error{
	db, err := sql.Open("sqlite3", "brew.db")
	if err != nil {
			fmt.Println("Error opening SQLite database:", err)
			return err
	}
  defer db.Close()

	var hashedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		fmt.Println("Error getting password from DB", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
			fmt.Println("Error checking password hash", err)
			return err
	}

	fmt.Println("User Login successful.")
	return nil
}