package core

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/000mrLuigi000/Library/logger"
)

//Event обьект событий
type Event struct {
	log *logger.Logger
}

var (
	KeyEnter = []byte{10, 0, 0, 0, 0}
	KeySpace = []byte{32, 0, 0, 0, 0}
)

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
			fmt.Println(reflect.DeepEqual(key, KeyEnter))
		}
	}
}
