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

	i := 1
	for {
		nc.Publish("intros", []byte(fmt.Sprintf("Hello World! - %d", i)))
		time.Sleep(1 * time.Second)
		fmt.Printf("Sent message %d\n", i)
		i++
	}
}

func getUrl() string {
	username := "developer"
	password := "a.very.secure.password!"
	hostport := "localhost:4222"

	return fmt.Sprintf("nats://%s:%s@%s", username, password, hostport)
}
