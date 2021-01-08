package core

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/000mrLuigi000/Library/logger"
)

//App главный обьект приложения
type App struct {
	w   *Window
	log *logger.Logger
	g   map[interface{}]interface{}
}

//InitApp инициализаци приложения
func InitApp() *App {
	a := &App{
		log: logger.InitLogger(""),
		g:   map[interface{}]interface{}{},
	}
	w := &Window{
		log: a.log,
	}
	a.w = w

	newLine, newColumn := w.size()
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
	clear()
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
	reset()
	return
}

//GetValue получить глобальное значение
func (a *App) GetValue(key interface{}) (interface{}, error) {
	rw.Lock()
	defer rw.Unlock()

	if v, ok := a.g[key]; ok {
		return v, nil
	}
	return nil, errors.New("Key not found")
}

//SetValue запомнить глобальное значение
func (a *App) SetValue(key, value interface{}) {
	rw.Lock()
	defer rw.Unlock()

	a.g[key] = value
	return
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

			time.Sleep(fct)
		}
	}
}
