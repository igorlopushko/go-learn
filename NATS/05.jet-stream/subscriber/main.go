package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	username string
	password string
	hostname = "localhost"
	port     = 4222
)

func init() {
	flag.StringVar(&username, "u", username, "username for NATS server")
	flag.StringVar(&password, "p", password, "password for NATS server")
	flag.StringVar(&hostname, "host", hostname, "hostname for NATS server")
	flag.IntVar(&port, "port", port, "port for NATS server")
	flag.Parse()
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := fmt.Sprintf("nats://%v:%v", hostname, port)
	if username != "" {
		url = fmt.Sprintf("nats://%v:%v@%v:%v", username, password, hostname, port)
	}

	nc, err := nats.Connect(url)
	fatalOnError(err)
	defer nc.Close()

	js, err := nc.JetStream()
	fatalOnError(err)

	// create ephemeral consumer
	//createEphemeralConsumer(js)

	// create durable consumer
	createDurableConsumer(js)

	log.Println("shutting down application...")
}

func createEphemeralConsumer(js nats.JetStreamContext) {
	js.Subscribe("orders.us", processMsg, nats.BindStream("ORDERS"))
	time.Sleep(10 * time.Second)
}

func createDurableConsumer(js nats.JetStreamContext) {
	_, err := js.AddConsumer("ORDERS", &nats.ConsumerConfig{
		Durable:      "my-consumer-1",
		Description:  "this is my awesome consumer",
		ReplayPolicy: nats.ReplayInstantPolicy,
	})
	fatalOnError(err)

	sub, err := js.PullSubscribe("orders.us", "my-consumer-1")
	fatalOnError(err)

	for sub.IsValid() {
		msgs, err := sub.Fetch(1)
		if err == nil {
			processMsg(msgs[0])
		}
	}

	time.Sleep(10 * time.Second)

	sub.Unsubscribe()
}

func processMsg(msg *nats.Msg) {
	fmt.Printf("INFO - Got reply - %s\n", string(msg.Data))
}
