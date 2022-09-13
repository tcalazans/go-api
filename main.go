package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Data string `json:data`
}

type swiftMusicData struct {
	Quote string `json:quote`
	Song  string `json:song`
	Album string `json:album`
}

func main() {
	router := gin.Default()
	router.GET("/myapi", func(c *gin.Context) {
		newData := data{Data: c.DefaultQuery("data", "")}
		// TODO: Ta concatenando
		c.IndentedJSON(http.StatusOK, newData)
		response, error := http.Get("https://taylorswiftapi.herokuapp.com/get")
		if error != nil {
			log.Fatalln(error)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
			c.JSON(200, gin.H{})
		}
	})

	router.Run("localhost:8080")
}
