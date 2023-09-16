package lifeBattle

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/eiannone/keyboard"
)

func Filler(my_map [50][50]string) [50][50]string {

	for m := 0; m < len(my_map); m++ {
		for n := 0; n < len(my_map); n++ {
			if my_map[m][n] == "" {
				my_map[m][n] = " "
			}
		}
	}

	return my_map
}

func plotter(my_map [50][50]string) {
	string_row := ""
	for m := 0; m < len(my_map); m++ {

		for n := 0; n < len(my_map); n++ {
			string_row = " " + string_row + my_map[m][n] + " "
		}

		fmt.Println(string_row)
		string_row = ""
	}
}

func Neighbour_count(my_map [50][50]string, m int, n int) int {
	counter := 0

	if my_map[m-1][n-1] == "■" {
		counter++
	}
	if my_map[m-1][n] == "■" {
		counter++
	}
	if my_map[m-1][n+1] == "■" {
		counter++
	}
	if my_map[m][n-1] == "■" {
		counter++
	}
	if my_map[m][n+1] == "■" {
		counter++
	}
	if my_map[m+1][n-1] == "■" {
		counter++
	}
	if my_map[m+1][n] == "■" {
		counter++
	}
	if my_map[m+1][n+1] == "■" {
		counter++
	}

	return counter
}

func Not_clear_field(my_map [50][50]string) bool {
	flag := false

	for m := 0; m < len(my_map); m++ {
		for n := 0; n < len(my_map); n++ {
			if my_map[m][n] == "■" {
				flag = true
			}
		}
	}
	return flag
}

func Gatling_celler(my_map [50][50]string, m int, n int) [50][50]string {

	return my_map
}

func Spawn_gliderITL(my_map [50][50]string, m int, n int) [50][50]string {

	my_map[m][n] = "■"
	my_map[m-1][n-1] = "■"
	my_map[m][n+1] = "■"
	my_map[m+1][n-1] = "■"
	my_map[m+1][n] = "■"

	return my_map
}

func Spawn_gliderUL(my_map [50][50]string, m int, n int) [50][50]string {

	my_map[m][n] = "■"
	my_map[m-1][n] = "■"
	my_map[m-1][n+1] = "■"
	my_map[m][n-1] = "■"
	my_map[m+1][n+1] = "■"

	return my_map
}

func Spawn_gliderDL(my_map [50][50]string, m int, n int) [50][50]string {

	my_map[m][n] = "■"
	my_map[m-1][n+1] = "■"
	my_map[m][n-1] = "■"
	my_map[m+1][n] = "■"
	my_map[m+1][n+1] = "■"

	return my_map
}

func Spawn_LWSSLR(my_map [50][50]string, m int, n int) [50][50]string {

	my_map[m][n] = "■"
	my_map[m-2][n] = "■"
	my_map[m-3][n+1] = "■"
	my_map[m-3][n+2] = "■"
	my_map[m-3][n+3] = "■"
	my_map[m-3][n+4] = "■"
	my_map[m-2][n+4] = "■"
	my_map[m-1][n+4] = "■"
	my_map[m][n+3] = "■"

	return my_map
}

func Spawn_LWSSRL(my_map [50][50]string, m int, n int) [50][50]string {

	my_map[m][n] = "■"
	my_map[m+1][n] = "■"
	my_map[m+2][n] = "■"
	my_map[m][n+1] = "■"
	my_map[m][n+2] = "■"
	my_map[m][n+3] = "■"
	my_map[m][n+4] = "■"
	my_map[m][n+5] = "■"
	my_map[m+1][n+6] = "■"
	my_map[m+3][n+6] = "■"
	my_map[m+4][n+4] = "■"
	my_map[m+4][n+3] = "■"
	my_map[m+3][n+1] = "■"

	return my_map
}

func SpawnRandFigure(my_map [50][50]string) [50][50]string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 4
	randInt := rand.Intn(max-min+1) + min
	minGL := 4
	maxGL := len(my_map) - 4
	randXY_GL1 := rand.Intn(maxGL-minGL+1) + minGL

	rand.Seed(time.Now().UnixNano())

	minLWSS := 8
	maxLWSS := len(my_map) - 8
	randXY_LWSS := rand.Intn(maxLWSS-minLWSS+1) + minLWSS

	switch randInt {
	case 0:
		my_map = Spawn_gliderITL(my_map, randXY_GL1, randXY_GL1)

	case 1:
		my_map = Spawn_gliderUL(my_map, randXY_GL1, randXY_GL1)

	case 2:
		my_map = Spawn_gliderDL(my_map, randXY_GL1, randXY_GL1)

	case 3:
		my_map = Spawn_LWSSLR(my_map, randXY_LWSS, randXY_LWSS)

	case 4:
		my_map = Spawn_LWSSRL(my_map, randXY_LWSS, randXY_LWSS)
	}

	return my_map
}

func SpawnRandFigureFromLeft(my_map [50][50]string) [50][50]string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1
	randInt := rand.Intn(max-min+1) + min
	minGL := 4
	maxGL := len(my_map) - 4
	randXY_GL1 := rand.Intn(maxGL-minGL+1) + minGL

	rand.Seed(time.Now().UnixNano())

	minLWSS := 8
	maxLWSS := len(my_map) - 8
	randXY_LWSS := rand.Intn(maxLWSS-minLWSS+1) + minLWSS

	switch randInt {
	case 0:
		my_map = Spawn_gliderITL(my_map, randXY_GL1, 4)

	case 1:
		my_map = Spawn_LWSSLR(my_map, randXY_LWSS, 8)
	}

	return my_map
}

func SpawnRandFigureFromRight(my_map [50][50]string) [50][50]string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2
	randInt := rand.Intn(max-min+1) + min
	minGL := 4
	maxGL := len(my_map) - 4
	randXY_GL1 := rand.Intn(maxGL-minGL+1) + minGL

	rand.Seed(time.Now().UnixNano())

	minLWSS := 8
	maxLWSS := len(my_map) - 8
	randXY_LWSS := rand.Intn(maxLWSS-minLWSS+1) + minLWSS

	switch randInt {
	case 0:
		my_map = Spawn_gliderUL(my_map, randXY_GL1, len(my_map)-4)

	case 1:
		my_map = Spawn_LWSSRL(my_map, randXY_LWSS, len(my_map)-8)

	case 2:
		my_map = Spawn_gliderDL(my_map, randXY_GL1, len(my_map)-4)

	}

	return my_map
}

func Keyboard_listener(space_pressed chan bool, nothing chan bool, interrupt chan os.Signal) {
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
			space_pressed <- true
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

func RenderNextGeneration(my_map [50][50]string /*, screen_map [50][50]string*/) [50][50]string {
	screen_map := my_map
	for m := 1; m < len(my_map)-1; m++ {
		for n := 1; n < len(my_map)-1; n++ {
			if my_map[m][n] == " " && Neighbour_count(my_map, m, n) == 3 {

				screen_map[m][n] = "■"

			}

			if (my_map[m][n] == "■") && (Neighbour_count(my_map, m, n) == 2 || Neighbour_count(my_map, m, n) == 3) {

				screen_map[m][n] = "■"

			}

			if (my_map[m][n] == "■") && !(Neighbour_count(my_map, m, n) == 2 || Neighbour_count(my_map, m, n) == 3) {

				screen_map[m][n] = " "

			}
		}
	}
	return screen_map
}

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SpawnStartGlider(my_map [50][50]string) [50][50]string {

	my_map[6][6] = "■"
	my_map[7][7] = "■"
	my_map[7][8] = "■"
	my_map[8][6] = "■"
	my_map[8][7] = "■"

	return my_map
}

func LifeBattle() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	space_pressed := make(chan bool)
	nothing := make(chan bool)
	go Keyboard_listener(space_pressed, nothing, interrupt)
	my_map := [50][50]string{}
	my_map = SpawnStartGlider(my_map)
	my_map = Filler(my_map)

	plotter(my_map)

	flag := true
	for flag {
		flag = Not_clear_field(my_map)

		select {

		case <-space_pressed:
			my_map = SpawnRandFigure(my_map)
			break

		case <-interrupt:
			return

		default:
			break

		}

		ClearTerminal()
		plotter(my_map)
		time.Sleep(70 * time.Millisecond)
		my_map = RenderNextGeneration(my_map)

	}

	return
}
