package server

import (
	life "net-life/life"

	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Recieve struct {
	space_pressed bool
}

type Message struct {
	Msg  string   `json:"msg"`
	Game game_map `json:"game"`
}

type game_map [50][50]string

var empty_map game_map

var msg Message

var upgrader = websocket.Upgrader{} // use default options

var connected_clients_count int = 0

var connections []*websocket.Conn

func startGame(c1 *websocket.Conn, c2 *websocket.Conn) {

	LifeBattle(c1, c2)

}

func recieve_loop(c *websocket.Conn, space_pressed chan bool, stop_game chan bool) {
	for {
		var recieve Recieve
		err := c.ReadJSON(&recieve)
		if err != nil {
			fmt.Println("Connection from client closed")
			stop_game <- true
			return
		}
		space_pressed <- true

	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	msg = Message{
		Msg:  "ok",
		Game: empty_map,
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	connected_clients_count++
	fmt.Println("Users count: ", connected_clients_count)
	connection := c
	connections = append(connections, connection)
	msg.Msg = "Welcome to the life. Wait for the second player..."

	c.WriteJSON(msg)

	if connected_clients_count == 2 {
		startGame(connections[0], connections[1])
	}

	if connected_clients_count > 2 {
		c.Close()
	}

}

func Server() {
	http.HandleFunc("/life", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func LifeBattle(c1 *websocket.Conn, c2 *websocket.Conn) {
	msg = Message{
		Msg:  "map",
		Game: empty_map,
	}
	space_pressed_from_c1 := make(chan bool)
	space_pressed_from_c2 := make(chan bool)
	stop_game := make(chan bool)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go recieve_loop(c1, space_pressed_from_c1, stop_game)
	go recieve_loop(c2, space_pressed_from_c2, stop_game)

	my_map := [50][50]string{}
	my_map = life.SpawnStartGlider(my_map)
	my_map = life.Filler(my_map)

	flag := true
	for flag {
		flag = life.Not_clear_field(my_map)

		select {
		case <-space_pressed_from_c1:
			my_map = life.SpawnRandFigureFromLeft(my_map)
			break
		case <-space_pressed_from_c2:
			my_map = life.SpawnRandFigureFromRight(my_map)
			break
		case <-stop_game:
			fmt.Println("Stop game!")
			c1.Close()
			c2.Close()
		case <-interrupt:
			msg.Msg = "server closed the game"
			c1.WriteJSON(msg)
			c2.WriteJSON(msg)
			os.Exit(0)
		default:
			break
		}

		time.Sleep(70 * time.Millisecond)
		my_map = life.RenderNextGeneration(my_map)
		msg.Game = my_map
		c1.WriteJSON(msg)
		c2.WriteJSON(msg)
	}

	return
}
