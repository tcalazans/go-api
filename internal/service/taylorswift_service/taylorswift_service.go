package taylorswift_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/tcalazans/go-api/internal/service"
)

type taylorSwiftService struct{}

func NewTaylorSwiftService() *taylorSwiftService {
	return &taylorSwiftService{}
}

func (s *taylorSwiftService) GetAlbum(ids []string) []*service.ApiResponse {
	var res []*service.ApiResponse
	var wg sync.WaitGroup
	wg.Add(len(ids))
	taylorQuotesChannel := make(chan *service.ApiResponse, len(ids))
	for i := 0; i < len(ids); i++ {
		go getTaylorQuotes(ids[i], taylorQuotesChannel, &wg)
	}
	wg.Wait()
	close(taylorQuotesChannel)
	for ch := range taylorQuotesChannel {
		res = append(res, ch)
	}
	return res
}

func getTaylorQuotes(newData string, channel chan *service.ApiResponse, wg *sync.WaitGroup) {
	defer Recover(channel)
	defer wg.Done()
	response, err := http.Get("https://taylorswiftapi.herokuapp.com/get?album=" + newData)
	apiResponse := &service.ApiResponse{}
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
	var swiftData service.SwiftMusicData
	json.Unmarshal(data, &swiftData)
	apiResponse.SwiftMusicData = &swiftData
	channel <- apiResponse
}

func Recover(c chan *service.ApiResponse) {
	if r := recover(); r != nil {
		fmt.Println("caiu no panic")
		c <- &service.ApiResponse{Partial: true}
	}
}
