// main.go (cmd package)

package main

import (
	"log"
	"net/http"

	handlers "github.com/sneha8080/userManagement/handler"
	repositories "github.com/sneha8080/userManagement/repository"
	"github.com/sneha8080/userManagement/services"
	"github.com/sneha8080/userManagement/utils"
)

func main() {
	// Load configuration from config.yaml
	// ...

	// Initialize the database connection
	db, err := utils.NewDB("mysql://username:password@localhost:3306/database_name")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Run database migrations
	err = db.Migrate()
	if err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db.DB)

	// Initialize services
	userService := services.NewUserService(userRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Register HTTP routes
	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/get", userHandler.GetUserByID)

	// Start the server
	log.Println("Server is running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
