A go interface to [CZMQ](http://czmq.zeromq.org)

This requires CZMQ head, and is targetted to be compatible with the next stable release of CZMQ.

Development is currently using CZMQ head compiled against ZeroMQ 4.0.4 Stable.

## Install

### Required

* ZeroMQ 4.0.4 or higher ( http://zeromq.org/intro:get-the-software )
* CZMQ Head ( https://github.com/zeromq/czmq )

### Get Go Library

  go get github.com/zeromq/goczmq

## Status

This library is alpha.  Not all features are complete.  API changes will happen.

Currently implemented:

* Sock
* Proxy
* Beacon
* Poller
* Auth

## Goals

Todo to finish inital phase::

* Gossip
* Loop
* Monitor

Secondary: Provide additional abstractions for "Go-isms" such as providing Zsocks as channel
accessable "services" within a go process.

## See Also

Peter Kleiweg's excellent zmq4 library for libzmq: http://github.com/pebbe/zmq4

## Smart Constructor Example
```go
package main

import (
	"flag"
	czmq "github.com/zeromq/goczmq"
	"log"
	"time"
)

func main() {
	var messageSize = flag.Int("message_size", 0, "size of message")
	var messageCount = flag.Int("message_count", 0, "number of messages")
	flag.Parse()

	pullSock, err := czmq.NewPULL("inproc://test")
	if err != nil {
		panic(err)
	}

	defer pullSock.Destroy()

	go func() {
		pushSock, err := czmq.NewPUSH("inproc://test")
		if err != nil {
			panic(err)
		}

		defer pushSock.Destroy()
		
		for i := 0; i < *messageCount; i++ {
			payload := make([]byte, *messageSize)
			err = pushSock.SendMessage(payload)
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
```

## Zbeacon Example
```go
package main

import (
	"fmt"
	czmq "github.com/taotetek/goczmq"
)

func main() {
	speaker := czmq.NewZbeacon()
	addr, err := speaker.Configure(9999)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Beacon configured on: %s\n", addr)

	listener := czmq.NewZbeacon()
	addr, err = listener.Configure(9999)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Beacon configured on: %s\n", addr)

	listener.Subscribe("HI")

	speaker.Publish("HI", 100)
	reply := listener.Recv(500)
	fmt.Printf("Received beacon: %v\n", reply)

	listener.Destroy()
	speaker.Destroy()
}
```

