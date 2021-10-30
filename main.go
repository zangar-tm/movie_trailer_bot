package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/zangar-tm/movie_trailer_bot/imdb"
)

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

func main() {
	botToken := "2020333417:AAESGYlh9fVVLPixWRXxjF2kFLPbtvaif8s"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("Something went wrong ", err.Error())
		}

		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	videoUrl, err := imdb.GetVideo(update.Message.Text)
	if err != nil {
		return err
	}
	botMessage.Text = videoUrl
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
