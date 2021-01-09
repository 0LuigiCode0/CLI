package core

import (
	"context"
	"os"

	"github.com/000mrLuigi000/Library/logger"
)

//Event обьект событий
type Event struct {
	log *logger.Logger
	windowI
}
type windowI interface {
	getLayout() LayoutI
}

//Key обьект кнопки
type Key struct {
	data []byte
	Name string
}

var (
	KeyEnter     = &Key{data: []byte{10, 0, 0}, Name: "Enter"}
	KeySpace     = &Key{data: []byte{32, 0, 0}, Name: "Space"}
	KeyBackSpace = &Key{data: []byte{127, 0, 0}, Name: "BackSpace"}
	KeyTab       = &Key{data: []byte{9, 0, 0}, Name: "Tab"}
	KeyEsc       = &Key{data: []byte{27, 0, 0}, Name: "Esc"}
	KeyUp        = &Key{data: []byte{27, 91, 65}, Name: "Up"}
	KeyDown      = &Key{data: []byte{27, 91, 66}, Name: "Down"}
	KeyLeft      = &Key{data: []byte{27, 91, 68}, Name: "Left"}
	KeyRight     = &Key{data: []byte{27, 91, 67}, Name: "Right"}
	Key0         = &Key{data: []byte{48, 0, 0}, Name: "0"}
	Key1         = &Key{data: []byte{49, 0, 0}, Name: "1"}
	Key2         = &Key{data: []byte{50, 0, 0}, Name: "2"}
	Key3         = &Key{data: []byte{51, 0, 0}, Name: "3"}
	Key4         = &Key{data: []byte{52, 0, 0}, Name: "4"}
	Key5         = &Key{data: []byte{53, 0, 0}, Name: "5"}
	Key6         = &Key{data: []byte{54, 0, 0}, Name: "6"}
	Key7         = &Key{data: []byte{55, 0, 0}, Name: "7"}
	Key8         = &Key{data: []byte{56, 0, 0}, Name: "8"}
	Key9         = &Key{data: []byte{57, 0, 0}, Name: "9"}
	KeyA         = &Key{data: []byte{97, 0, 0}, Name: "a"}
	KeyB         = &Key{data: []byte{98, 0, 0}, Name: "b"}
	KeyC         = &Key{data: []byte{99, 0, 0}, Name: "c"}
	KeyD         = &Key{data: []byte{100, 0, 0}, Name: "d"}
	KeyE         = &Key{data: []byte{101, 0, 0}, Name: "e"}
	KeyF         = &Key{data: []byte{102, 0, 0}, Name: "f"}
	KeyG         = &Key{data: []byte{103, 0, 0}, Name: "g"}
	KeyH         = &Key{data: []byte{104, 0, 0}, Name: "h"}
	KeyI         = &Key{data: []byte{105, 0, 0}, Name: "i"}
	KeyJ         = &Key{data: []byte{106, 0, 0}, Name: "j"}
	KeyK         = &Key{data: []byte{107, 0, 0}, Name: "k"}
	KeyL         = &Key{data: []byte{108, 0, 0}, Name: "l"}
	KeyM         = &Key{data: []byte{109, 0, 0}, Name: "m"}
	KeyN         = &Key{data: []byte{110, 0, 0}, Name: "n"}
	KeyO         = &Key{data: []byte{111, 0, 0}, Name: "o"}
	KeyP         = &Key{data: []byte{112, 0, 0}, Name: "p"}
	KeyQ         = &Key{data: []byte{113, 0, 0}, Name: "q"}
	KeyR         = &Key{data: []byte{114, 0, 0}, Name: "r"}
	KeyS         = &Key{data: []byte{115, 0, 0}, Name: "s"}
	KeyT         = &Key{data: []byte{116, 0, 0}, Name: "t"}
	KeyU         = &Key{data: []byte{117, 0, 0}, Name: "u"}
	KeyV         = &Key{data: []byte{118, 0, 0}, Name: "v"}
	KeyW         = &Key{data: []byte{119, 0, 0}, Name: "w"}
	KeyX         = &Key{data: []byte{120, 0, 0}, Name: "x"}
	KeyY         = &Key{data: []byte{121, 0, 0}, Name: "y"}
	KeyZ         = &Key{data: []byte{122, 0, 0}, Name: "z"}
	KeyTilda     = &Key{data: []byte{96, 0, 0}, Name: "~"}
	KeyPlus      = &Key{data: []byte{43, 0, 0}, Name: "+"}
	KeyMinus     = &Key{data: []byte{45, 0, 0}, Name: "-"}
	KeyEqual     = &Key{data: []byte{61, 0, 0}, Name: "="}
)
var keyList = []*Key{
	KeyEnter,
	KeySpace,
	KeyTab,
	KeyEsc,
	KeyUp,
	KeyDown,
	KeyLeft,
	KeyRight,
	Key0,
	Key1,
	Key2,
	Key3,
	Key4,
	Key5,
	Key6,
	Key7,
	Key8,
	Key9,
	KeyA,
	KeyB,
	KeyC,
	KeyD,
	KeyE,
	KeyF,
	KeyG,
	KeyH,
	KeyI,
	KeyJ,
	KeyK,
	KeyL,
	KeyM,
	KeyN,
	KeyO,
	KeyP,
	KeyQ,
	KeyR,
	KeyS,
	KeyT,
	KeyU,
	KeyV,
	KeyW,
	KeyX,
	KeyY,
	KeyZ,
	KeyTilda,
	KeyPlus,
	KeyMinus,
}

func (e *Event) listen(ctx context.Context) {
	defer wg.Done()
	c := make(chan []byte)
	var b, key []byte = make([]byte, 3), make([]byte, 3)
	go func() {
		for {
			os.Stdin.Read(b)
			c <- b
			b = make([]byte, 3)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			//e.log.Info("Listen event stoped")
			return
		case key = <-c:
			e.findKey(key)
		}
	}
}

func (e *Event) findKey(key []byte) {
	for _, k := range keyList {
		if equals(k.data, key) {
			if l := e.windowI.getLayout(); l != nil {
				if f, err := l.getEvent(k); err == nil {
					f()
				}
			}
			return
		}
	}

	return
}
