package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tcalazans/go-api/internal/server/controller"
)

func GetServer(controller controller.SwiftController) *gin.Engine {
	router := gin.Default()
	router.GET("/myapi", controller.GetAlbum)
	return router
}
