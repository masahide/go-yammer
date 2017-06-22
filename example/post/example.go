package main

import (
	"encoding/json"
	"flag"
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
	var to int
	var mes string
	flag.IntVar(&to, "to", 0, "replied to id:(threadId)")
	flag.StringVar(&mes, "mes", "", "messages")
	flag.Parse()
	if to == 0 {
		log.Fatal("required \"to\" flag")
	}
	conf := loadCache("cache.json")
	client := yammer.New(conf.AccessToken)
	postMes := &yammer.CreateMessageParams{Body: mes, RepliedToId: to}
	_, err := client.PostMessage(postMes)
	log.Println(err)
}
