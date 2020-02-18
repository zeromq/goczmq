package main

import (
	"flag"
	"log"

	czmq "github.com/zeromq/goczmq/v4"
)

func main() {
	bindTo := flag.String("bind_to", "", "address to bind to")
	messageSize := flag.Int("message_size", 0, "size of message")
	roundtripCount := flag.Int("roundtrip_count", 0, "number of roundtrips")
	flag.Parse()

	if *bindTo == "" || *messageSize == 0 || *roundtripCount == 0 {
		log.Fatalln("usage: local_lat <bind-to> <message-size> <roundtrip-count>")
	}

	// Create socket.
	repSock, err := czmq.NewRep(*bindTo)
	if err != nil {
		log.Fatalf("Failed to create REP socket: %s", err)
	}
	defer repSock.Destroy()

	// Receive messages and reply.
	for i := 0; i != *roundtripCount; i++ {
		msg, err := repSock.RecvMessage()
		if err != nil {
			log.Fatalf("Failed to receive message: %s", err)
		}

		if len(msg) != 1 {
			log.Fatalf("Message of incorrect size received: %d", len(msg))
		}

		if len(msg[0]) != *messageSize {
			log.Fatalf("Message of incorrect size received: %d", len(msg[0]))
		}

		err = repSock.SendMessage(msg)
		if err != nil {
			log.Fatalf("Failed to send message: %s", err)
		}
	}
}
