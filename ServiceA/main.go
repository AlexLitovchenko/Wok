package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

func main() {

	url := "https://www.cbr-xml-daily.ru/latest.js"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	defer nc.Close()

	js.Publish("foo", data)

}
