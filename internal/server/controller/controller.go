package controller

import "github.com/gin-gonic/gin"

type SwiftController interface {
	GetAlbum(context *gin.Context)
}
