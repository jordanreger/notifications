package main

import (
	"time"
)

type MastodonAccount struct {
	ID          int    `json:"id"`
	Username    string `json:"acct"`
	DisplayName string `json:"display_name"`
}

type MastodonStatus struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	Content   string     `json:"content"`
}

type MastodonNotification struct {
	ID        int             `json:"id"`
	Type      string          `json:"type"`
	CreatedAt *time.Time      `json:"created_at"`
	Account   MastodonAccount `json:"account"`
	Status    MastodonStatus  `json:"status"`
}
