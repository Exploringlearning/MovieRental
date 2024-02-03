package main

import (
	"bufio"
	"log"
	"movierental/internal/routes"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	loadEnv()
	engine := gin.Default()
	routes.RegisterRoutes(engine)
	err := engine.Run("localhost:8080")
	if err != nil {
		return
	}
}

func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error opening .env file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			err := os.Setenv(key, value)
			if err != nil {
				log.Fatal("Error while setting env variables:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading .env file:", err)
	}
}
