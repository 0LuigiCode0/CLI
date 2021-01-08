package core

import (
	"context"
	"fmt"
	"os"

	"github.com/000mrLuigi000/Library/logger"
)

//Event обьект событий
type Event struct {
	log *logger.Logger
}

//Key обьект кнопки
type Key struct {
	data []byte
	Name string
}

var (
	KeyEnter = &Key{data: []byte{10, 0, 0, 0, 0}, Name: "Enter"}
	//KeySpace = &Key{data: []byte{32, 0, 0, 0, 0}, Name: "Space"}
	KeyTab   = &Key{data: []byte{9, 0, 0, 0, 0}, Name: "Tab"}
	KeyEsc   = &Key{data: []byte{27, 0, 0, 0, 0}, Name: "Esc"}
	KeyUp    = &Key{data: []byte{27, 91, 65, 0, 0}, Name: "Up"}
	KeyDown  = &Key{data: []byte{27, 91, 66, 0, 0}, Name: "Down"}
	KeyLeft  = &Key{data: []byte{27, 91, 68, 0, 0}, Name: "Left"}
	KeyRight = &Key{data: []byte{27, 91, 67, 0, 0}, Name: "Right"}
)
var KeyList = []*Key{
	KeyEnter,
	//KeySpace,
	KeyTab,
	KeyEsc,
	KeyUp,
	KeyDown,
	KeyLeft,
	KeyRight,
}

func (e *Event) listen(ctx context.Context) {
	defer wg.Done()
	c := make(chan []byte)
	var b, key []byte = make([]byte, 5), make([]byte, 5)
	go func() {
		for {
			os.Stdin.Read(b)
			c <- b
			b = make([]byte, 5)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			//e.log.Info("Listen event stoped")
			return
		case key = <-c:
			fmt.Println(findKey(key))
		}
	}
}
