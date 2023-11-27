package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/igorlopushko/go-learn/NATS/02.request-reply/model"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	nc.Subscribe("intros", func(msg *nats.Msg) {
		var data model.Payload
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Fatal("can't unmarshal incoming request")
		}
		fmt.Printf("I got a message: %v\n", data)

		replyData := fmt.Sprintf("ack message # %v", data.Count)
		err = msg.Respond([]byte(replyData))
		if err != nil {
			log.Fatal("can't respond for the incoming request")
		}

	})

	time.Sleep(1 * time.Hour)
}
