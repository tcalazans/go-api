package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Data string
}

type swiftMusicData struct {
	Quote string `json:quote`
	Song  string `json:song`
	Album string `json:album`
}

func main() {
	router := gin.Default()
	router.GET("/myapi", func(c *gin.Context) {
		newData, exist := c.GetQuery("data")
		if exist {
			c.JSON(http.StatusOK, data{Data: newData})
		} else {
			getTaylorQuotes(c)
		}
	})
	router.Run("localhost:8080")
}

func getTaylorQuotes(c *gin.Context) {
	response, error := http.Get("https://taylorswiftapi.herokuapp.com/get")
	if error != nil {
		log.Fatalln(error)
	}
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	var swiftData swiftMusicData
	json.Unmarshal(data, &swiftData)
	c.JSON(http.StatusOK, swiftData)
}
