package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/Serg0/nats-websocket-gw"
)

func usage() {
	fmt.Printf(`Usage: %s [ --help ] [ --no-origin-check ] [ --trace ]
`, os.Args[0])
}

func main() {

    fmt.Printf(`HELLO %s`, '')
    fmt.Printf(`NATS_HOST: %s `, os.Getenv("NATS_HOST"))
	settings := gw.Settings{
		NatsAddr: os.Getenv("NATS_HOST"),
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--help":
			usage()
			return
		case "--no-origin-check":
			settings.WSUpgrader = &websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
				CheckOrigin:     func(r *http.Request) bool { return true },
			}
		case "--trace":
			settings.Trace = true
		default:
			fmt.Printf("Invalid args: %s\n\n", arg)
			usage()
			return
		}
	}

	gateway := gw.NewGateway(settings)
	http.HandleFunc("/nats", gateway.Handler)
	http.ListenAndServe("0.0.0.0:8910", nil)
}
