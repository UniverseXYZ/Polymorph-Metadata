package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/polymorph-metadata/handlers"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	funcframework.RegisterHTTPFunction("/token", handlers.TokenMetadata)

	// Use PORT environment variable, or default to 8080.
	port := "9090"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
