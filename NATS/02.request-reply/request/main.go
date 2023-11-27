package main

import (
	"encoding/json"
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

	count := 0
	for {
		payload := &model.Payload{
			Data:  "Hello World!",
			Count: count,
		}

		data, _ := json.Marshal(payload)

		reply, err := nc.Request("intros", []byte(data), 500*time.Millisecond)
		if err != nil {
			log.Printf("error sending message count = %v, err: %v", count, err)
			continue
		}

		time.Sleep(1 * time.Second)
		count++
		log.Printf("sent message %v, got reply %v", count, string(reply.Data))
	}
}
