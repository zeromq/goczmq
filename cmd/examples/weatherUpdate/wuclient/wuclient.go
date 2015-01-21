//
// Weather proxy listens to weather server which is constantly
// emitting weather data
// Binds SUB socket to tcp://127.0.0.1:5556
//
// Usage : go run examples/wuclient.go 87654

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	czmq "github.com/zeromq/goczmq"
)

func main() {
	pubEndpoint := "tcp://127.0.0.1:5556"
	totalTemperature := 0

	filter := "59937"
	if len(os.Args) > 1 {
		filter = string(os.Args[1])
	}

	subSock, err := czmq.NewSUB(pubEndpoint, filter)
	if err != nil {
		panic(err)
	}

	defer subSock.Destroy()

	fmt.Printf("Collecting updates from weather server for %s…\n", filter)
	subSock.Connect(pubEndpoint)

	for i := 0; i < 100; i++ {
		msg, _, err := subSock.RecvFrame()
		if err != nil {
			panic(err)
		}

		weatherData := strings.Split(string(msg), " ")
		temperature, err := strconv.ParseInt(weatherData[1], 10, 64)
		if err == nil {
			totalTemperature += int(temperature)
		}
	}

	fmt.Printf("Average temperature for zipcode %s was %dF \n\n", filter,
		totalTemperature/100)
}
