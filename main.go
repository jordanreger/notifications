package main

import "time"

var bsky_host = "https://bsky.social"
var bsky_username = "jordanreger.com"

var mastodon_host = "https://mastodon.sdf.org"

var interval = 5 * time.Minute

func main() {
	for {
		<-time.After(interval)
		GetBskyNotifications()
		GetMastodonNotifications()
	}
}
