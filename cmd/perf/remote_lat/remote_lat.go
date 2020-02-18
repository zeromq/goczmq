package main

import (
	"flag"
	"log"
	"time"

	czmq "github.com/zeromq/goczmq/v4"
)

func main() {
	connectTo := flag.String("connect_to", "", "address to bind to")
	messageSize := flag.Int("message_size", 0, "size of message")
	roundtripCount := flag.Int("roundtrip_count", 0, "number of roundtrips")
	flag.Parse()

	if *connectTo == "" || *messageSize == 0 || *roundtripCount == 0 {
		log.Fatalln("usage: remote_lat <connect-to> <message-size> <roundtrip-count>")
	}

	// Create socket.
	reqSock, err := czmq.NewReq(*connectTo)
	if err != nil {
		log.Fatalf("Failed to create REQ socket: %s", err)
	}
	defer reqSock.Destroy()

	// Create message.
	msg := [][]byte{make([]byte, *messageSize)}

	// Start timing.
	startTime := time.Now()

	// Send messages and read replies.
	for i := 0; i != *roundtripCount; i++ {
		err := reqSock.SendMessage(msg)
		if err != nil {
			log.Fatalf("Failed to send message: %s", err)
		}

		reply, err := reqSock.RecvMessage()
		if err != nil {
			log.Fatalf("Failed to receive message: %s", err)
		}

		if len(reply) != 1 {
			log.Fatalf("Message of incorrect size received: %d", len(reply))
		}

		if len(reply[0]) != *messageSize {
			log.Fatalf("Message of incorrect size received: %d", len(reply[0]))
		}
	}

	duration := time.Since(startTime)
	microseconds := duration.Seconds() * 1e6
	latency := microseconds / float64((*roundtripCount * 2))

	log.Printf("message size: %d [B]\n", *messageSize)
	log.Printf("roundtrip count: %d\n", *roundtripCount)
	log.Printf("average latency: %.3f [us]\n", latency)
}
