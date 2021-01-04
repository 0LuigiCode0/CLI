package core

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/000mrLuigi000/Library/logger"
)

var rw sync.Mutex
var wg sync.WaitGroup

//App главный обьект приложения
type App struct {
	w   *window
	fct time.Duration
	log *logger.Logger
}

func game(App *App) {
	close := make(chan os.Signal)
	signal.Notify(close, os.Interrupt, os.Kill)

	for {
		select {
		case <-close:
			return
		default:
			i := int(rand.Float32()*100) % App.w.lines
			j := int(rand.Float32()*1000) % App.w.column
			x := byte(int(rand.Float32()*100)%App.w.lines + 50)
			App.w.setPX(i, j, fmt.Sprintf("\033[5m\033[48;5;%vm\033[38;5;%vm%v\033[0m", x, x+20, string(x)))

			time.Sleep(App.fct)
		}
	}
}

//InitApp инициализаци приложения
func InitApp() *App {
	a := &App{
		fct: time.Millisecond * 16,
		log: logger.InitLogger(""),
	}
	w := &window{
		log:      a.log,
		apptowin: a,
		fct:      a.fct,
	}
	a.w = w

	newLine, newColumn := a.size()
	frame := make([][]string, newLine, newLine)
	for i := range frame {
		frame[i] = make([]string, newColumn, newColumn)
	}
	w.lines = newLine
	w.column = newColumn
	w.frame = frame

	return a
}

//Start запуск приложения
func (a *App) Start() {
	a.clear()
	wg.Add(2)

	ctx, cancelf := context.WithCancel(context.Background())

	go a.w.reView(ctx)
	go a.w.reSize(ctx)
	go game(a)

	close := make(chan os.Signal)
	signal.Notify(close, os.Interrupt, os.Kill)
	<-close
	cancelf()
	wg.Wait()
	a.reset()
	return
}

func (a *App) clear() {
	a.cmd(true, "clear")
	a.cmd(true, "tput", "civis")
	a.cmd(true, "stty", "-echo")
}
func (a *App) reset() {
	a.cmd(true, "clear")
	a.cmd(true, "tput", "reset")
	a.cmd(true, "stty", "echo")
}

func (a *App) size() (int, int) {
	var newLine, newColumn int
	res, _ := a.cmd(false, "tput", "lines")
	if len(res) > 0 {
		newLine, _ = strconv.Atoi(string(res[:len(res)-1]))
	}
	res, _ = a.cmd(false, "tput", "cols")
	if len(res) > 0 {
		newColumn, _ = strconv.Atoi(string(res[:len(res)-1]))
	}
	return newLine, newColumn
}

func (a *App) cmd(out bool, comand string, args ...string) ([]byte, error) {
	v := exec.Command(comand, args...)
	v.Stdin = os.Stdin
	if out {
		v.Stdout = os.Stdout
		v.Run()
	}
	return v.Output()
}
