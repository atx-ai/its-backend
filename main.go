package main

import (
	"log"
	"net/http"

	"github.com/atx-ai/its-backend/controller"
	localdb "github.com/atx-ai/its-backend/db"
	_ "github.com/atx-ai/its-backend/docs" // Import the swag auto-generated docs file
	"github.com/atx-ai/its-backend/model"
	"github.com/atx-ai/its-backend/service"
	"github.com/go-chi/chi/middleware"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func main() {
	// Initialize database connection
	var err error
	dbConn, err := connectDB() // Assign to the global db variable
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	err = dbConn.AutoMigrate(&model.Issue{})
	if err != nil {
		log.Fatalf("failed to auto migrate schema: %v", err)
	}

	// Auto migrate the schema
	err = dbConn.AutoMigrate(&model.Commnet{})
	if err != nil {
		log.Fatalf("failed to auto migrate schema: %v", err)
	}

	// Initialize service and controller
	issueService := service.NewIssueService(dbConn)
	commentService := service.NewCommnetService(dbConn)
	issueController := controller.NewIssueController(issueService)
	commnetController := controller.NewCommnetController(commentService)

	// Initialize chi router
	router := chi.NewRouter()

	// Middleware setup
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		MaxAge:         1800,
	}))

	// Swagger documentation endpoint
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	router.Mount("/issues", issueController.Routes())
	router.Mount("/issues/{issueID}/comments", commnetController.Routes())

	// Start HTTP server
	log.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}

func connectDB() (*gorm.DB, error) {
	// Define database connection options
	dbOptions := localdb.DBOptions{
		Username: "issue_tracker",
		Password: "issue_tracker",
		Host:     "localhost",
		Port:     "5432",
		DBName:   "issue_tracker",
		SSLMode:  "disable",
		TimeZone: "UTC",
	}

	// Connect to the database
	return localdb.ConnectDB(dbOptions)
}
