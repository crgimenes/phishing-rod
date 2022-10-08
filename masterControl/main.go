package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	origin, ok := os.LookupEnv("MASTERCONTROL_ORIGIN")
	if !ok {
		log.Fatal("origin not set")
	}

	url, ok := os.LookupEnv("MASTERCONTROL_URL")
	if !ok {
		log.Fatal("url not set")
	}

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("{\"SUBSCRIBE\": \"TIMESTAMP\"}\n")); err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write([]byte("{\"SUBSCRIBE\": \"GPS_LOCATION\"}\n")); err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write([]byte("{\"SUBSCRIBE\": \"MESSAGE\"}\n")); err != nil {
		log.Fatal(err)
	}

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
	}
}
