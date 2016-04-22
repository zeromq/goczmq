//
// Weather update server
// Binds PUB socket to tcp://127.0.0.1:5556
// Publishes random weather updates
//

package main

import (
	"fmt"
	"math/rand"

	czmq "gopkg.in/zeromq/goczmq.v1"
)

func main() {
	pubEndpoint := "tcp://127.0.0.1:5556"
	pubSock, err := czmq.NewPub(pubEndpoint)
	if err != nil {
		panic(err)
	}

	defer pubSock.Destroy()
	pubSock.Bind(pubEndpoint)

	for {
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 85
		relHumidity := rand.Intn(50) + 10

		msg := fmt.Sprintf("%d %d %d", zipcode, temperature, relHumidity)
		err := pubSock.SendFrame([]byte(msg), 0)
		if err != nil {
			panic(err)
		}
	}
}
