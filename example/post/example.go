package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/masahide/go-yammer/yammer"
)

type cache struct {
	AccessToken string
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
	postMes := &yammer.CreateMessageParams{Body: "hoge", RepliedToId: 11111}
	_, err := client.PostMessage(postMes)
	log.Println(err)
}
