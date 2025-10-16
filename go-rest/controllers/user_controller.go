package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/manish-npx/go-lang/go-rest/models"
)

const CONTENT_TYPE = "Content-Type"
const APPLICATION_JSON = "application/json"

// user slice having userData
var users = []models.User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// get user  w=write r=request
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	// Encode (convert) the users slice into JSON and send
	json.NewEncoder(w).Encode(users)
}

// GetUserByID handles GET /user?id=1
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get the "id" query parameter from URL
	idStr := r.URL.Query().Get("id")
	// Convert the string to an integer
	id, _ := strconv.Atoi(idStr)

	// Search for the user by ID
	for _, user := range users {
		if user.ID == id {
			w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// If not found, return a 404 error
	http.Error(w, "User not found", http.StatusNotFound)
}

// CreateUser handles POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Create a variable to hold the new user data from the request body
	var newUser models.User

	// Decode the JSON body into our newUser struct
	json.NewDecoder(r.Body).Decode(&newUser)

	// Assign a new ID (auto-increment style)
	newUser.ID = len(users) + 1
	// Add to the in-memory slice
	users = append(users, newUser)

	// Return the newly created user as JSON
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	json.NewEncoder(w).Encode(newUser)
}
