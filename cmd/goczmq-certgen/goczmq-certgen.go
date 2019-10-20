package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zeromq/goczmq/v4"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("certificate filename required")
	}

	filename := args[0]

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading name")
	}

	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading email")
	}

	fmt.Print("Organization: ")
	organization, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading organization")
	}

	fmt.Print("Version: ")
	version, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading version")
	}

	cert := goczmq.NewCert()
	cert.SetMeta("name", strings.TrimSpace(name))
	cert.SetMeta("email", strings.TrimSpace(email))
	cert.SetMeta("organization", strings.TrimSpace(organization))
	cert.SetMeta("version", strings.TrimSpace(version))
	cert.Save(filename)
}
