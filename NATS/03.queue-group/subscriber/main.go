package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/igorlopushko/go-learn/NATS/03.queue-group/model"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	sub, err := nc.QueueSubscribe("intros", "queue1", processMsg)
	if err != nil {
		log.Fatalf("can't subscribe to NATS queue 'queue1': %v", err)
	}
	defer sub.Unsubscribe()

	time.Sleep(1 * time.Hour)
}

func processMsg(msg *nats.Msg) {
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
}
