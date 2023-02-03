package main

import (
	"test-task-rlm/pkg/routes"
	"test-task-rlm/pkg/services"
)

func main() {
	services.ParseParams()
	routes.SetupRoutes()
}
