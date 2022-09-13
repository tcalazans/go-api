package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiResponse struct {
	SwiftMusicData *swiftMusicData `json:"swift_music_data,omitempty"`
	Partial        bool            `json:"partial"`
}

type swiftMusicData struct {
	Quote string `json:"quote,omitempty"`
	Song  string `json:"song,omitempty"`
	Album string `json:"album,omitempty"`
}

func main() {
	router := gin.Default()
	router.GET("/myapi", func(c *gin.Context) {
		newData, _ := c.GetQuery("album")
		res, err := getTaylorQuotes(c, newData)
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
