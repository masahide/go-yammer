package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/masahide/go-yammer/cometd"
	"github.com/masahide/go-yammer/yammer"
)

type cache struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	Extra        interface{}
}

func loadCache(file string) cache {
	var conf cache
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&conf); err != nil {
		log.Fatal(err)
	}
	return conf
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf := loadCache("cache.json")
	client := yammer.New(conf.AccessToken)

	realtime, err := client.Realtime()
	if err != nil {
		log.Println(err)
		return
	}
	inbox, err := client.InboxFeedV2()
	if err != nil {
		log.Println(err)
		return
	}

	rt := cometd.New(realtime.RealtimeURI, realtime.AuthenticationToken)
	err = rt.Handshake()
	if err != nil {
		log.Println(err)
		return
	}

	rt.SubscribeToFeed(inbox.ChannelID)
	messageChan := make(chan *cometd.ConnectionResponse, 10)
	//messageChan := make(chan *schema.MessageFeed, 10)
	stopChan := make(chan bool, 1)

	log.Printf("Polling Realtime channelID: %v\n", inbox.ChannelID)
	go rt.Poll(messageChan, stopChan)
	for {
		select {
		case m := <-messageChan:
			if m.Channel == "/meta/connect" {
				continue
			}
			if m.Data.Type != "message" {
				log.Printf("Data.Type is not message. channel:%#v", m)
				continue
			}
			for _, mes := range m.Data.Feed.Messages {
				log.Println(mes.SenderId, mes.Body.Parsed)
			}
		}
	}
}
