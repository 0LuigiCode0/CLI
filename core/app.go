package core

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"github.com/000mrLuigi000/Library/logger"
)

//App главный обьект приложения
type App interface {
	Start()
	GetValue(key interface{}) (interface{}, error)
	SetValue(key, value interface{})
	Window() Window
}

type app struct {
	w   Window
	e   *Event
	log *logger.Logger
	g   map[interface{}]interface{}
}

//InitApp инициализаци приложения
func InitApp(layout ILayout) (App, error) {
	log := logger.InitLogger("")
	a := &app{
		log: log,
		g:   map[interface{}]interface{}{},
	}
	if layout == nil {
		return nil, errors.New("Layout is nil")
	}
	w := &window{
		log:    log,
		layout: layout,
	}
	a.w = w
	e := &Event{
		log:     log,
		windowI: w,
	}
	a.e = e

	newLine, newColumn := w.size()
	frame := make([][]string, newLine, newLine)
	for i := range frame {
		frame[i] = make([]string, newColumn, newColumn)
	}
	w.lines = newLine
	w.column = newColumn
	w.frame = frame

	return a, nil
}

//Start запуск приложения
func (a *app) Start() {
	clear()
	wg.Add(1)

	ctx, cancelf := context.WithCancel(context.Background())
	//	go a.w.reView(ctx)
	//go a.w.reSize(ctx)
	go a.e.listen(ctx)

	if f := a.w.getLayout().getCreate(); f != nil {
		f()
	}

	close := make(chan os.Signal)
	signal.Notify(close, os.Interrupt, os.Kill)
	<-close

	if l := a.w.getLayout(); l != nil {
		if f := l.getDelete(); f != nil {
			f()
		}
	}

	cancelf()
	wg.Wait()
	reset()
	return
}

//GetValue получить обьект окна
func (a *app) Window() Window {
	return a.w
}

//GetValue получить глобальное значение
func (a *app) GetValue(key interface{}) (interface{}, error) {
	rw.Lock()
	defer rw.Unlock()

	if v, ok := a.g[key]; ok {
		return v, nil
	}
	return nil, errors.New("Key not found")
}

//SetValue запомнить глобальное значение
func (a *app) SetValue(key, value interface{}) {
	rw.Lock()
	defer rw.Unlock()

	a.g[key] = value
	return
}

// func game(App *app) {
// 	close := make(chan os.Signal)
// 	signal.Notify(close, os.Interrupt, os.Kill)

// 	for {
// 		select {
// 		case <-close:
// 			return
// 		default:
// 			i := int(rand.Float32()*100) % App.w.lines
// 			j := int(rand.Float32()*1000) % App.w.column
// 			x := byte(int(rand.Float32()*100)%App.w.lines + 50)
// 			App.w.setPX(i, j, fmt.Sprintf("\033[5m\033[48;5;%vm\033[38;5;%vm%v\033[0m", x, x+20, string(x)))

// 			time.Sleep(fct)
// 		}
// 	}
// }
