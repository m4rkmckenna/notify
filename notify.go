package main

import (
	"flag"
	"net/http"
	"bytes"
	"encoding/json"
	"os"
)

const (
	pushbullet_base_uri string = "https://api.pushbullet.com/v2/pushes"
)

func main() {
	var token string
	var title string
	var body string
	flag.StringVar(&token, "token", "", "Pushbullet access token")
	flag.StringVar(&title, "title", "", "Notification title")
	flag.StringVar(&body, "body", "", "Notification body")
	flag.Parse()
	pushNote(token, title, body)
	os.Exit(0)
}


func pushNote(token string, title string, body string) {
	jsonBody, _ := json.Marshal(Push{Title:title, Body:body, Type:"note"})
	req, _ := http.NewRequest("POST", pushbullet_base_uri, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", token)
	http.DefaultClient.Do(req)
}

type Push struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Type  string `json:"type"`
}
