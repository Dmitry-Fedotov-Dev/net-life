package client

import (
	"fmt"
	life "net-life/life"
	"net/url"
	"os"
	"os/signal"

	"github.com/eiannone/keyboard"
	"github.com/gorilla/websocket"
)

type Message struct {
	Msg  string   `json:"msg"`
	Game game_map `json:"game"`
}

type SpacePressed struct {
	Space bool `json:"space"`
}

type game_map [50][50]string

var msg Message

var empty_map game_map

func recieve_loop(c *websocket.Conn, rcv chan Message, game chan game_map) {
	for {
		var msg Message

		err := c.ReadJSON(&msg)
		if err != nil {
			c.Close()
		}

		if msg.Msg == "map" {
			game <- msg.Game
		}
		rcv <- msg
	}

}

func keyboard_listener(c *websocket.Conn, interrupt chan os.Signal) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents

		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeySpace {
			var space_pressed = SpacePressed{
				Space: true,
			}
			c.WriteJSON(space_pressed)
		}

		if event.Key == keyboard.KeyEsc {
			break
		}

		if event.Key == keyboard.KeyCtrlC {
			interrupt <- os.Interrupt
			break
		}

	}
}

func Client() {
	rcv := make(chan Message)
	game := make(chan game_map)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/life"}
	fmt.Println("connecting to: ", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dial:", err)
	}
	defer c.Close()

	go recieve_loop(c, rcv, game)
	go keyboard_listener(c, interrupt)

	for {
		select {

		case recieve := <-game:
			life.ClearTerminal()
			plotter(recieve)

		case recieve := <-rcv:
			fmt.Println(recieve.Msg)

		case <-interrupt:
			fmt.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write close:", err)
				return
			}
			return
		}
	}
}

func plotter(my_map game_map) {
	string_row := ""
	for m := 0; m < len(my_map); m++ {

		for n := 0; n < len(my_map); n++ {
			string_row = " " + string_row + my_map[m][n] + " "
		}

		fmt.Println(string_row)
		string_row = ""
	}
}
