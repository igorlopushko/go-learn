package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	url := getUrl()

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	nc.Subscribe("intros", func(msg *nats.Msg) {
		fmt.Printf("I got a message: %s\n", string(msg.Data))
	})

	time.Sleep(1 * time.Hour)
}

func getUrl() string {
	username := "developer"
	password := "a.very.secure.password!"
	hostport := "localhost:4222"

	return fmt.Sprintf("nats://%s:%s@%s", username, password, hostport)
}
