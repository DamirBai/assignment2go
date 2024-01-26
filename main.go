package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Email   string
	Address string
	Age     int
}

var db *gorm.DB

func InitializeDatabase() {
	dsn := "user=postgres dbname=auth password=admin sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
}
func registrationHTMLHandler(w http.ResponseWriter, r *http.Request) {
	renderHTMLPage(w, "registration.html")
}
func GetUserHTMLHandler(w http.ResponseWriter, r *http.Request) {
	renderHTMLPage(w, "getuser.html")
}

func UserHTMLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}
	renderGetHTMLPage(w, "user.html", user)
}
func DeleteUserHTMLHandler(w http.ResponseWriter, r *http.Request) {
	renderHTMLPage(w, "delete_user.html")
}
func UpdateNameHTMLHandler(w http.ResponseWriter, r *http.Request) {
	renderHTMLPage(w, "update_name.html")
}

func renderHTMLPage(w http.ResponseWriter, page string) {
	tmpl, err := template.ParseFiles(page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	tmpl.Execute(w, nil)
}
func renderGetHTMLPage(w http.ResponseWriter, page string, user User) {
	tmpl, err := template.ParseFiles(page)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, user)
}
func UpdateNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	userID := vars["id"]
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	existingUser, err := getUserByID(userID)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}
	existingUser.Name = updatedUser.Name

	if err := db.Save(&existingUser).Error; err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "200", "message": "User updated successfully"})
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]

	if err := db.Delete(&User{}, userID).Error; err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "200", "message": "User deleted successfully"})
}
func main() {
	InitializeDatabase()
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/registration", registrationHTMLHandler).Methods("GET")
	r.HandleFunc("/get_user", GetUserHTMLHandler).Methods("GET")
	r.HandleFunc("/user/{id}", UserHTMLHandler).Methods("GET")
	r.HandleFunc("/update_name", UpdateNameHTMLHandler).Methods("GET")
	r.HandleFunc("/delete_user", DeleteUserHTMLHandler).Methods("Get")

	r.HandleFunc("/create_user", CreateUserHandler).Methods("POST")
	r.HandleFunc("/delete_user/{id}", DeleteUserHandler).Methods("POST")
	r.HandleFunc("/update_name/{id}", UpdateNameHandler).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newUser).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "201", "message": "User created successfully"})
}
func getUserByID(userID string) (User, error) {
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
