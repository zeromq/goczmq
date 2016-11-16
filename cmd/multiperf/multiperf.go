package main

import (
	"flag"
	"log"
	"time"

	czmq "github.com/zeromq/goczmq"
)

func main() {
	var messageSize = flag.Int("message_size", 0, "size of message")
	var messageCount = flag.Int("message_count", 0, "number of messages")
	flag.Parse()

	pullSock, err := czmq.NewPull("inproc://test")
	if err != nil {
		panic(err)
	}

	defer pullSock.Destroy()

	go func() {
		pushSock, err := czmq.NewPush("inproc://test")
		if err != nil {
			panic(err)
		}

		defer pushSock.Destroy()
		for i := 0; i < *messageCount; i++ {
			payload := make([]byte, *messageSize)
			err = pushSock.SendMessage([][]byte{payload})
			if err != nil {
				panic(err)
			}
		}
	}()

	startTime := time.Now()
	for i := 0; i < *messageCount; i++ {
		msg, err := pullSock.RecvMessage()
		if err != nil {
			panic(err)
		}
		if len(msg) != 1 {
			panic("msg too small")
		}
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	throughput := float64(*messageCount) / elapsed.Seconds()
	megabits := float64(throughput*float64(*messageSize)*8.0) / 1e6

	log.Printf("message size: %d", *messageSize)
	log.Printf("message count: %d", *messageCount)
	log.Printf("test time (seconds): %f", elapsed.Seconds())
	log.Printf("mean throughput: %f [msg/s]", throughput)
	log.Printf("mean throughput: %f [Mb/s]", megabits)
}
