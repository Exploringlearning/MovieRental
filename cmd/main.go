package main

import (
	"movierental/internal/routes"
)

func main() {
	engine := routes.InitRouter()
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
