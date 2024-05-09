package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type mastodon_n_res []MastodonNotification

func ParseMastodonPost(txt string) string {
	doc, err := html.Parse(strings.NewReader(txt))
	if err != nil {
		fmt.Println(err)
	}

	res := ""

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			res += n.Data
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return res
}

func GetMastodonNotifications() {
	godotenv.Load(".env")
	client := &http.Client{}

	token := os.Getenv("mastodontoken")

	req, err := http.NewRequest("GET", mastodon_host+"/api/v1/notifications", bytes.NewBufferString(""))
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var n_body mastodon_n_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &n_body)

	notifications := n_body
	for _, notification := range notifications {
		if time.Now().Sub(*notification.CreatedAt).Minutes() <= interval.Minutes() && notification.Type == "mention" {

			user := notification.Account.DisplayName
			handle := notification.Account.Username
			message := ParseMastodonPost(notification.Status.Content)
			date := notification.CreatedAt

			SendMastodonNotification(user, handle, message, date)
		}
	}

	s_req, err := http.NewRequest("POST", mastodon_host+"/api/v1/notifications/clear", bytes.NewBufferString(""))
	s_req.Header.Add("Authorization", "Bearer "+token)
	res, err = client.Do(s_req)
	if err != nil {
		fmt.Println(err)
	}
}
