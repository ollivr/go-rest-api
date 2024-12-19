package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful-api/internal/routes"
)

type User struct {
	Id              string
	Name            string
	Label           string
	Slug            string
	Email           string
	Phone           string
	Website         string
	Post            string
	Location        string
	AccountName     string
	MustVerifyEmail bool
	EmailVerifiedAt string
	Enabled         bool
	CreatedAt       string
	UpdatedAt       string
}

const EndpointHitMessage = "Endpoint Hit: %s"

func AllUsers(w http.ResponseWriter, r *http.Request) {
	initializeDB := routes.InitializeDB
	db, err := initializeDB() // Initialize the database connection
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var users []User // Fetch all users from the database
	if err := db.Find(&users).Error; err != nil {
		log.Println("Error fetching users from the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil { // Encode users to JSON and send response
		log.Println("Error encoding users to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Fetched users: %+v\n", users) // Debug: Log the fetched users (optional, remove if unnecessary)
}

func logEndpointHit(endpoint string) { // Utility function to log endpoint hits
	fmt.Printf(EndpointHitMessage+"\n", endpoint)
}
func HandleNewUser(w http.ResponseWriter, r *http.Request) { // Renamed functions for clarity and a consistent naming convention
	logEndpointHit("newUser")
}
func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	logEndpointHit("deleteUser")
}
func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	logEndpointHit("updateUser")
}
func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	logEndpointHit("homePage")
}
