package service

type ApiResponse struct {
	// padrao snakecase em go na parte do Json
	SwiftMusicData *SwiftMusicData `json:"swift_music_data,omitempty"`
	Partial        bool            `json:"partial"`
}

type SwiftMusicData struct {
	// omitempty vai omitir caso o objeto esteja vazio
	Quote string `json:"quote,omitempty"`
	Song  string `json:"song,omitempty"`
	Album string `json:"album,omitempty"`
}

type SwiftService interface {
	GetAlbum(splitedData []string) []*ApiResponse
}
