package main

import (
	"log"
	"net/http"

	"github.com/atx-ai/its-backend/controller"
	localdb "github.com/atx-ai/its-backend/db"
	_ "github.com/atx-ai/its-backend/docs" // Import the swag auto-generated docs file
	"github.com/atx-ai/its-backend/model"
	"github.com/atx-ai/its-backend/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	// Initialize service and controller
	issueService := service.NewIssueService(dbConn)
	issueController := controller.NewIssueController(issueService)

	// Initialize chi router
	router := chi.NewRouter()

	// Middleware setup
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Routes setup
	router.Route("/issues", func(r chi.Router) {
		r.Get("/{id}", issueController.GetIssue)
		r.Post("/", issueController.CreateIssue)
		r.Put("/{id}", issueController.UpdateIssue)
		r.Delete("/{id}", issueController.DeleteIssue)
		r.Get("/", issueController.ListIssues)
	})

	// Swagger documentation endpoint
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

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
