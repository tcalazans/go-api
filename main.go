package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

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
		var wg sync.WaitGroup
		newData, _ := c.GetQuery("album")
		arrayData := strings.Split(newData, ",")
		wg.Add(len(arrayData))
		taylorQuotesChannel := make(chan *apiResponse, len(arrayData))
		for i := 0; i < len(arrayData); i++ {
			fmt.Println(arrayData)
			go getTaylorQuotes(c, arrayData[i], taylorQuotesChannel, &wg)
		}
		wg.Wait()
		close(taylorQuotesChannel)
		for ch := range taylorQuotesChannel {
			res = append(res, ch)
			fmt.Println(ch)
		}
		for _, r := range res {
			if r.Partial {
				c.JSON(http.StatusPartialContent, res)
				return
			}
		}
		c.JSON(http.StatusOK, res)
	})
	router.Run("localhost:8080")
}

func getTaylorQuotes(c *gin.Context, newData string, channel chan *apiResponse, wg *sync.WaitGroup) {
	defer Recover(channel)
	defer wg.Done()
	response, err := http.Get("https://taylorswiftapi.herokuapp.com/get?album=" + newData)
	apiResponse := &apiResponse{}
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if string(data) == "" {
		panic("no data found")
	}
	var swiftData swiftMusicData
	json.Unmarshal(data, &swiftData)
	apiResponse.SwiftMusicData = &swiftData
	channel <- apiResponse
}

func Recover(c chan *apiResponse) {
	if r := recover(); r != nil {
		fmt.Println("caiu no panic")
		c <- &apiResponse{Partial: true}
	}
}
