package main

import (
	"github.com/tcalazans/go-api/internal/server"
	"github.com/tcalazans/go-api/internal/server/controller/swift_controller"
	"github.com/tcalazans/go-api/internal/service/swift_service"
)

func main() {
	swiftService := swift_service.NewSwiftService()
	swiftController := swift_controller.NewSwiftController(swiftService)
	router := server.GetServer(swiftController)
	router.Run("localhost:8080")
}
