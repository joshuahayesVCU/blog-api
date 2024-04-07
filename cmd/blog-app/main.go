package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joshuahayesVCU/blog-api/pkg/config"
	"github.com/joshuahayesVCU/blog-api/pkg/model"
)

func main() {
	// Initialize configuration
	err := config.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := model.InitDb(os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	setupRoutes(router, db)
	router.Run(":8080")

}

func setupRoutes(router *gin.Engine, db *model.DB) {

}
