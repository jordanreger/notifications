package main

import (
	"encoding/base64"
	"fmt"
	//"github.com/joho/godotenv"
	"log"
	"mime"
	"net/mail"
	"net/smtp"
	"os"
	"time"
)

func SendBskyNotification(user string, handle string, message string, date *time.Time) {
	//godotenv.Load(".env")
	auth := smtp.PlainAuth("", os.Getenv("smtpuser"), os.Getenv("smtppass"), "smtp.mail.me.com")

	from := mail.Address{Name: "Bluesky", Address: "notifications@jordanreger.com"}
	to := mail.Address{Name: "Jordan Reger", Address: "mail@jordanreger.com"}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = mime.QEncoding.Encode("UTF-8", "New notification from "+user)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	body := user + " (@" + handle + ")\n\n" +
		message + "\n\n\n" +
		date.Format("1/2/2006 at 15:04 UTC")

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail("smtp.mail.me.com:587", auth, from.Address, []string{to.Address}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func SendMastodonNotification(user string, handle string, message string, date *time.Time) {
	//godotenv.Load(".env")
	auth := smtp.PlainAuth("", os.Getenv("smtpuser"), os.Getenv("smtppass"), "smtp.mail.me.com")

	from := mail.Address{Name: "Mastodon", Address: "notifications@jordanreger.com"}
	to := mail.Address{Name: "Jordan Reger", Address: "mail@jordanreger.com"}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = mime.QEncoding.Encode("UTF-8", "New notification from "+user)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	body := user + " (@" + handle + ")\n\n" +
		message + "\n\n\n" +
		date.Format("1/2/2006 at 15:04 UTC")

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail("smtp.mail.me.com:587", auth, from.Address, []string{to.Address}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
