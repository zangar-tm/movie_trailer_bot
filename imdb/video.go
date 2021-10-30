package imdb

import (
	"net/http"
)

const (
	apiKey     = "k_oukuscby"
	getDataUrl = "https://imdb-api.com/en/API/SearchTitle/"
	videoUrl   = "https://www.imdb.com/video/"
)

type DataResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Id string `json:"id"`
}

func retrieveVideo() (Result, error) {

}

func makeRequest(getDataUrl string) (*http.Request, error) {

}
