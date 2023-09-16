package main

import (
	client "net-life/client"
	life "net-life/life"
	server "net-life/server"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "mp":
			multiPlayer()
		case "server":
			serverMode()
		default:
			singlePlayer()
		}
	} else {
		singlePlayer()
	}

}

func singlePlayer() {
	life.LifeBattle()
}

func multiPlayer() {
	client.Client()
}

func serverMode() {
	server.Server()
}
