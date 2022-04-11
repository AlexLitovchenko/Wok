package handler

import (
	"log"

	"github.com/nats-io/nats.go"
)

var Sub *nats.Subscription

func Connection() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	Sub, err = nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}
	//defer nc.Close()
}
