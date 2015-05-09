# goczmq [![Build Status](https://travis-ci.org/zeromq/goczmq.svg?branch=master)](https://travis-ci.org/zeromq/goczmq) [![Doc Status](https://godoc.org/github.com/zeromq/goczmq?status.png)](https://godoc.org/github.com/zeromq/goczmq)

## Introduction
A golang interface to [CZMQ](http://czmq.zeromq.org)

This currently requires CZMQ head, and is targetted to be compatible with the upcoming 3.x release of CZMQ.

## Installation

```
git clone git@github.com:jedisct1/libsodium.git
cd libsodium
./autogen.sh; ./configure; make; make check
sudo make install
sudo ldconfig
```

```
git clone git@github.com:zeromq/libzmq.git
cd libzmq
./autogen.sh; ./configure --with-libsodium; make; make check
sudo make install
sudo ldconfig
```

```
git clone git@github.com:zeromq/czmq.git
cd czmq
./autogen.sh; ./configure; make; make check
sudo make install
sudo ldconfig
```

```
go get github.com/zeromq/goczmq
```

## Usage
```go
package main

import (
	"time"
	"github.com/zeromq/goczmq"
)

func main() {
	payload := []byte("Hello")
	payload_size := len(payload)
	count := 1000000
	
	pullSock, err := goczmq.NewPULL("inproc://test")
	if err != nil {
		panic(err)
	}

	defer pullSock.Destroy()

	go func() {
		pushSock, err := goczmq.NewPUSH("inproc://test")
		if err != nil {
			panic(err)
		}

		defer pushSock.Destroy()
		
		for i := 0; i < count; i++ {
			err = pushSock.SendFrame([]byte("Hello"), 0)
			if err != nil {
				panic(err)
			}
		}
	}()

	startTime := time.Now()

	for i := 0; i < count; i++ {
		msg, err := pullSock.RecvMessage()
		if err != nil {
			panic(err)
		}
		if string(msg[0]) != "Hello" {
			panic("invalid msg!")
		}
	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	throughput := float64(count) / elapsed.Seconds()
	megabits := float64(throughput*float64(payload_size)*8.0) / 1e6

	log.Printf("message size: %d", payload_size)
	log.Printf("message count: %d", count)
	log.Printf("test time (seconds): %f", elapsed.Seconds())
	log.Printf("mean throughput: %f [msg/s]", throughput)
	log.Printf("mean throughput: %f [Mb/s]", megabits)
}
```

## GoDoc
[godoc](https://godoc.org/github.com/zeromq/goczmq)

## License
This project uses the MPL v2 license, see LICENSE
