package main

import (
	"github.com/tcalazans/go-api/internal/server"
	"github.com/tcalazans/go-api/internal/server/controller/taylorswift_controller"
	"github.com/tcalazans/go-api/internal/service/taylorswift_service"
)

type apiResponse struct {
	// padrao snakecase em go na parte do Json
	SwiftMusicData *swiftMusicData `json:"swift_music_data,omitempty"`
	Partial        bool            `json:"partial"`
}

type swiftMusicData struct {
	// omitempty vai omitir caso o objeto esteja vazio
	Quote string `json:"quote,omitempty"`
	Song  string `json:"song,omitempty"`
	Album string `json:"album,omitempty"`
}

func main() {
	taylorSwiftService := taylorswift_service.NewTaylorSwiftService()
	taylorSwiftController := taylorswift_controller.NewTaylorSwiftController(taylorSwiftService)
	router := server.GetServer(taylorSwiftController)
	router.Run("localhost:8080")
}
