package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"time"
)

func GetJWT(username string, password string) BskySession {
	client := &http.Client{}

	body, err := json.Marshal(&BskyUser{Identifier: username, Password: password})
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", bsky_host+"/xrpc/com.atproto.server.createSession", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var session BskySession
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &session)
	defer res.Body.Close()

	return session
}

type bsky_n_res struct {
	Notifications []BskyNotification `json:"notifications,omitempty"`
}

func GetBskyNotifications() {
	client := &http.Client{}

	godotenv.Load(".env")

	jwt := GetJWT(bsky_username, os.Getenv("bskypass")).AccessJWT

	req, err := http.NewRequest("GET", bsky_host+"/xrpc/app.bsky.notification.listNotifications", bytes.NewBufferString(""))
	req.Header.Add("Authorization", "Bearer "+jwt)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var n_body bsky_n_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &n_body)

	notifications := n_body.Notifications

	for _, notification := range notifications {
		if !notification.IsRead && notification.Record.Type == "app.bsky.feed.post" {
			user := notification.Author.DisplayName
			handle := notification.Author.Handle
			message := notification.Record.Text
			date := notification.Record.CreatedAt

			SendBskyNotification(user, handle, message, date)
		}
	}

	type s_body struct {
		SeenAt time.Time `json:"seenAt"`
	}

	s_res, _ := json.Marshal(&s_body{SeenAt: time.Now()})

	s_req, err := http.NewRequest("POST", bsky_host+"/xrpc/app.bsky.notification.updateSeen", bytes.NewBuffer(s_res))
	s_req.Header.Add("Content-Type", "application/json")
	s_req.Header.Add("Authorization", "Bearer "+jwt)
	res, err = client.Do(s_req)
	if err != nil {
		fmt.Println(err)
	}

}
