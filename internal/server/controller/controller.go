package controller

import "github.com/gin-gonic/gin"

type TaylorSwiftController interface {
	GetAlbum(c *gin.Context)
}
