package main

import (
	bsky "jordanreger.com/bsky/api"
	"time"
)

type BskyUser struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type BskySession struct {
	AccessJWT  string `json:"accessJwt"`
	RefreshJWT string `json:"refreshJwt"`
	Handle     string `json:"handle"`
	DID        string `json:"did"`
}

type BskyNotification struct {
	URI       string      `json:"uri"`
	CID       string      `json:"cid"`
	Author    bsky.Actor  `json:"author"`
	Reason    string      `json:"reason"`
	Record    bsky.Record `json:"record"`
	IsRead    bool        `json:"isRead"`
	IndexedAt time.Time   `json:"indexedAt"`
}
