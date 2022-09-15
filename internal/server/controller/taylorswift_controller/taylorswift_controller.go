package taylorswift_controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/tcalazans/go-api/internal/service"
)

type taylorSwiftController struct {
	taylorSwiftService service.TaylorSwiftService
}

func NewTaylorSwiftController(service service.TaylorSwiftService) *taylorSwiftController {
	return &taylorSwiftController{
		taylorSwiftService: service,
	}
}

func (s *taylorSwiftController) GetAlbum(c *gin.Context) {
	newData, _ := c.GetQuery("album")
	ids := strings.Split(newData, ",")
	res := s.taylorSwiftService.GetAlbum(ids)
	for _, v := range res {
		if v.Partial {
			c.JSON(http.StatusPartialContent, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}
