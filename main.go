package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/myapi", func(c *gin.Context) {
		var res []*apiResponse
		var result *apiResponse
		var err error
		newData, _ := c.GetQuery("album")
		arrayData := strings.Split(newData, ",")
		for i := 0; i < len(arrayData); i++ {
			result, err = getTaylorQuotes(c, arrayData[i])
			res = append(res, result)
		}
		if err != nil {
			c.JSON(http.StatusPartialContent, res)
			return
		}
		c.JSON(http.StatusOK, res)
	})
	router.Run("localhost:8080")
}

func getTaylorQuotes(c *gin.Context, newData string) (*apiResponse, error) {
	response, err := http.Get("https://taylorswiftapi.herokuapp.com/get?album=" + newData)
	apiResponse := &apiResponse{}
	if err != nil {
		apiResponse.Partial = true
		return apiResponse, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		apiResponse.Partial = true
		return apiResponse, err
	}
	if string(data) == "" {
		return apiResponse, nil
	}
	var swiftData swiftMusicData
	json.Unmarshal(data, &swiftData)
	apiResponse.SwiftMusicData = &swiftData
	return apiResponse, nil
}
