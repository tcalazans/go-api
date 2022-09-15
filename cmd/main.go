package main

import (
	"github.com/tcalazans/go-api/internal/server"
	"github.com/tcalazans/go-api/internal/server/controller/taylorswift_controller"
	"github.com/tcalazans/go-api/internal/service/taylorswift_service"
)

func main() {
	taylorSwiftService := taylorswift_service.NewTaylorSwiftService()
	taylorSwiftController := taylorswift_controller.NewTaylorSwiftController(taylorSwiftService)
	router := server.GetServer(taylorSwiftController)
	router.Run("localhost:8080")
}
