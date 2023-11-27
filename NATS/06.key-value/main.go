package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

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

	listBuckets(js)

	// create a new bucket called 'sensor'
	bucketName := "sensor"
	sensonrs, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: bucketName,
	})
	fatalOnError(err)

	// add a few key/value pairs
	sensonrs.PutString("temperature", "48deg")
	sensonrs.PutString("humidity", "50%")
	sensonrs.PutString("pressure", "10bars")

	listBuckets(js)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("shutting down application...")
}

func listBuckets(js nats.JetStreamContext) {
	// list all buckets all the keys
	for bucketName := range js.KeyValueStoreNames() {
		// trim first 3 characters from bucketName
		bucketName = bucketName[3:]

		fmt.Printf("Bucket - %s\n", bucketName)

		bucket, err := js.KeyValue(bucketName)
		if err != nil {
			log.Printf("failed to bind a bucket: %v", err)
			continue
		}

		keys, err := bucket.Keys()
		if err != nil {
			continue
		}

		for _, key := range keys {
			val, _ := bucket.Get(key)
			fmt.Printf("\t%s - %s\n", key, val.Value())
		}
	}
}
