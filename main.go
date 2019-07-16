package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "Tu token de telegram",
		URL:    "https://api.telegram.org",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	//Responde un mensaje
	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})
	//Manda un audio
	b.Handle("/musicr", func(m *tb.Message) {
		a := &tb.Audio{File: tb.FromDisk("nombre_archivo.mp3")}
		b.Send(m.Sender, a)
	})
	//Manda una imagen
	b.Handle("/imagen", func(m *tb.Message) {
		ai := &tb.Photo{File: tb.FromDisk("akamaru.png")}
		b.Send(m.Sender, ai)
	})

	b.Start()
}
