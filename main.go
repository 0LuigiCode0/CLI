package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

//App главный обьект приложения
var App = newApp()

type app struct {
	w     *window
	ctx   map[string]interface{}
	fps   time.Duration
	close chan os.Signal

	rw sync.Mutex
}

type window struct {
	lines  int
	column int
	layout [][]string

	rw sync.Mutex
}

func newApp() *app {
	a := &app{
		ctx:   map[string]interface{}{},
		fps:   time.Millisecond * 100,
		close: make(chan os.Signal),
	}
	signal.Notify(a.close, os.Interrupt, os.Kill)
	w := &window{}
	a.w = w

	newLine, newColumn := a.size()
	frame := make([][]string, newLine, newLine)
	for i := range frame {
		frame[i] = make([]string, newColumn, newColumn)
	}
	w.lines = newLine
	w.column = newColumn
	w.layout = frame

	return a
}

func (w *window) reSize() {
	for {
		select {
		case <-App.close:
			return
		default:
			newLine, newColumn := App.size()
			if newLine != w.lines && newColumn != w.column {
				App.clear()
				frame := make([][]string, newLine, newLine)
				for i := range frame {
					frame[i] = make([]string, newColumn, newColumn)
				}
				w.lines = newLine
				w.column = newColumn
				w.setLayout(frame)
			} else if newLine != w.lines {
				App.clear()
				frame := make([][]string, newLine, newLine)
				for i := range frame {
					frame[i] = make([]string, w.column, w.column)
				}
				w.lines = newLine
				w.setLayout(frame)
			} else if newColumn != w.column {
				App.clear()
				frame := w.layout
				for i := range frame {
					frame[i] = make([]string, newColumn, newColumn)
				}
				w.column = newColumn
				w.setLayout(frame)
			}

			time.Sleep(App.fps)
		}
	}
}

func (w *window) reView() {
	for {
		select {
		case <-App.close:
			return
		default:
			resp := "\033[0;0H"
			for i := 0; i < w.lines; i++ {
				for j := 0; j < w.column; j++ {
					//px := w.getPX(i, j)
					resp += fmt.Sprintf("\033[5m\033[48;5;%vm\033[38;5;%vmH\033[0m", 20, 20+20)
					// if px == "" {
					// 	resp += " "
					// } else {
					// 	resp += px
					// }
				}
				if i < w.lines-1 {
					resp += "\n"
				}
				fmt.Print(resp)
				resp = ""
			}

			time.Sleep(App.fps)
		}
	}
}

func (w *window) reBuild() {
	for {
		select {
		case <-App.close:
			return
		default:
			for i := 0; i < w.lines; i++ {
				for j := 0; j < w.column; j++ {
					w.setPX(i, j, w.layout[i][j])
				}
			}

			time.Sleep(App.fps)
		}
	}
}

func (a *app) clear() {
	root := exec.Command("clear")
	root.Stdout = os.Stdout
	root.Run()
}

func (a *app) size() (int, int) {
	var newLine, newColumn int
	res, _ := a.Cmd("tput", "lines")
	if len(res) > 0 {
		newLine, _ = strconv.Atoi(string(res[:len(res)-1]))
	}
	res, _ = a.Cmd("tput", "cols")
	if len(res) > 0 {
		newColumn, _ = strconv.Atoi(string(res[:len(res)-1]))
	}
	return newLine, newColumn
}

func (a *app) Cmd(comand string, args ...string) ([]byte, error) {
	v := exec.Command(comand, args...)
	return v.CombinedOutput()
}

func (a *app) getByCtx(key string) interface{} {
	a.rw.Lock()
	defer a.rw.Unlock()

	return a.ctx[key]
}

func (a *app) setByCtx(key string, value interface{}) {
	a.rw.Lock()
	defer a.rw.Unlock()

	a.ctx[key] = value
	return
}

func (w *window) getPX(i, j int) string {
	w.rw.Lock()
	defer w.rw.Unlock()

	return w.layout[i][j]
}

func (w *window) setPX(i, j int, v string) {
	w.rw.Lock()
	defer w.rw.Unlock()

	w.layout[i][j] = v
	return
}

func (w *window) setLayout(arr [][]string) {
	w.rw.Lock()
	defer w.rw.Unlock()

	w.layout = arr
	fmt.Print(len(arr))
	return
}

func main() {
	defer App.clear()
	defer App.Cmd("tput", "reset")
	App.Cmd("tput", "civis")
	App.clear()

	go App.w.reView()
	go App.w.reSize()
	go App.w.reBuild()
	go game()

	<-App.close
	os.Stdout.Close()
	return
}

func game() {
	for {
		select {
		case <-App.close:
			return
		default:
			i := int(rand.Float32()*100) % App.w.lines
			j := int(rand.Float32()*1000) % App.w.column
			x := byte(int(rand.Float32()*100)%App.w.lines + 50)
			App.w.setPX(i, j, fmt.Sprintf("\033[5m\033[48;5;%vm\033[38;5;%vm%v\033[0m", x, x+20, string(x)))

			time.Sleep(App.fps)
		}
	}
}
