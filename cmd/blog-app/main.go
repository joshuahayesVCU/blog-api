package main

import (
	"log"
	"os"

	"github.com/joshuahayesVCU/blog_api/config"
)

func main() {
	// Initialize configuration
	err := config.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v")
	}

	Sprintf(os.Getenv("PORT"))
) 
