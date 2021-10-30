package imdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	getMovieUrl   = "https://imdb-api.com/en/API/SearchTitle/"
	getTrailerUrl = "https://imdb-api.com/en/API/Trailer/"
	videoUrl      = "https://www.imdb.com/video/"
)

type RestResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Id string `json:"id"`
}

type Video struct {
	VideoId string `json:"videoId"`
}

func GetVideo(inputData string) (string, error) {
	trailerId, err := makeTrailerRequest(inputData)
	if err != nil {
		return "", err
	}
	return videoUrl + trailerId, nil
}

func makeTrailerRequest(inputData string) (string, error) {
	videoId, err := getVideoId(inputData)
	if err != nil {
		return "", err
	}
	url := getTrailerUrl + os.Getenv("apiKey") + "/" + videoId
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var video Video
	err = json.Unmarshal(body, &video)
	if err != nil {
		return "", err
	}
	return video.VideoId, nil
}

func getVideoId(inputData string) (string, error) {
	results, err := makeVideoRequest(inputData)
	if err != nil {
		return "", err
	}
	if len(results) < 1 {
		return "", errors.New("Nothing found")
	}
	return results[0].Id, nil
}

func makeVideoRequest(inputData string) ([]Result, error) {
	url := getMovieUrl + os.Getenv("apiKey") + "/" + inputData
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Results, nil
}
