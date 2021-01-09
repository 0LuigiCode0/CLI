package core

import (
	"errors"
	"fmt"
)

//LayoutI интерфейс слоя
type LayoutI interface {
	AddEvent(key *Key, f func()) error
	getEvent(key *Key) (func(), error)
}
type layout struct {
	event  map[*Key]func()
	blocks []*Block
}

//NewLayout создание нового слоя
func NewLayout() LayoutI {
	lay := &layout{
		event: map[*Key]func(){},
	}

	lay.event[KeyTab] = lay.onTab
	lay.event[KeyEnter] = lay.onEnter
	lay.event[KeyBackSpace] = lay.onBackSpace
	lay.event[KeyUp] = lay.onUp
	lay.event[KeyDown] = lay.onDown
	lay.event[KeyLeft] = lay.onUp
	lay.event[KeyRight] = lay.onDown

	return lay
}

func (l *layout) AddEvent(key *Key, f func()) error {
	rw.Lock()
	defer rw.Unlock()

	if key == KeyTab || key == KeyEnter || key == KeyBackSpace || key == KeyUp || key == KeyDown || key == KeyLeft || key == KeyRight {
		return fmt.Errorf("System key [ %v ] cannot be set", key.Name)
	}
	l.event[key] = f
	return nil
}

func (l *layout) getEvent(key *Key) (func(), error) {
	rw.Lock()
	defer rw.Unlock()

	if f, ok := l.event[key]; ok {
		return f, nil
	}
	return nil, errors.New("Key not found")
}

func (l *layout) onTab() {

}

func (l *layout) onEnter() {
	fmt.Println("on enter")
}

func (l *layout) onBackSpace() {

}

func (l *layout) onUp() {

}

func (l *layout) onDown() {

}

func (l *layout) onLeft() {

}

func (l *layout) onRigth() {

}
