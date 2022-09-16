package swift_controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tcalazans/go-api/internal/service"
)

type SwiftController struct {
	swiftService service.SwiftService
}

func NewSwiftController(service service.SwiftService) *SwiftController {
	return &SwiftController{
		swiftService: service,
	}
}

func (s *SwiftController) GetAlbum(c *gin.Context) {
	newData, _ := c.GetQuery("album")
	splitedData := strings.Split(newData, ",")
	res := s.swiftService.GetAlbum(splitedData)
	for _, v := range res {
		if v.Partial {
			c.JSON(http.StatusPartialContent, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}
