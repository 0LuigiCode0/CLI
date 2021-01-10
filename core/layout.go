package core

import (
	"errors"
	"fmt"
)

//ILayout интерфейс слоя
type ILayout interface {
	AddEvent(key *Key, f func()) error
	OnCreate(f func()) ILayout
	OnUpdate(f func()) ILayout
	OnDelete(f func()) ILayout
	SetStyle(style map[Style]interface{}) ILayout
	SetComponents(comp ...IComponent) ILayout
	getCreate() func()
	getDelete() func()
	getEvent(key *Key) (func(), error)
	getStyle() map[Style]interface{}
	getComponents() []IComponent
}

type layout struct {
	event      map[*Key]func()
	components []IComponent
	style      map[Style]interface{}
	onCreate   func()
	onUpdate   func()
	onDelete   func()
}

//Layout создание нового слоя
func Layout() ILayout {
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

func (l *layout) OnCreate(f func()) ILayout {
	rw.Lock()
	defer rw.Unlock()

	l.onCreate = f
	return l
}

func (l *layout) OnUpdate(f func()) ILayout {
	rw.Lock()
	defer rw.Unlock()

	l.onUpdate = f
	return l
}

func (l *layout) OnDelete(f func()) ILayout {
	rw.Lock()
	defer rw.Unlock()

	l.onDelete = f
	return l
}

func (l *layout) SetStyle(style map[Style]interface{}) ILayout {
	rw.Lock()
	defer rw.Unlock()

	l.style = style
	if f := l.onUpdate; f != nil {
		f()
	}
	return l
}

func (l *layout) SetComponents(comp ...IComponent) ILayout {
	rw.Lock()
	defer rw.Unlock()

	l.components = comp
	if f := l.onUpdate; f != nil {
		f()
	}
	return l
}

func (l *layout) getCreate() func() {
	rw.Lock()
	defer rw.Unlock()

	return l.onCreate
}

func (l *layout) getDelete() func() {
	rw.Lock()
	defer rw.Unlock()

	return l.onDelete
}

func (l *layout) getEvent(key *Key) (func(), error) {
	rw.Lock()
	defer rw.Unlock()

	if f, ok := l.event[key]; ok {
		return f, nil
	}
	return nil, errors.New("Key not found")
}

func (l *layout) getStyle() map[Style]interface{} {
	rw.Lock()
	defer rw.Unlock()

	return l.style
}

func (l *layout) getComponents() []IComponent {
	rw.Lock()
	defer rw.Unlock()

	return l.components
}

func (l *layout) onTab() {
	for _, c := range l.getComponents() {
		if f := c.getOnTab(); c.GetActive() && f != nil {
			f()
		}
	}
}

func (l *layout) onEnter() {
	for _, c := range l.getComponents() {
		if f := c.getOnEnter(); c.GetActive() && f != nil {
			f()
		}
	}
}

func (l *layout) onBackSpace() {
	for _, c := range l.getComponents() {
		if f := c.getOnBackSpace(); c.GetActive() && f != nil {
			f()
		}
	}
}

func (l *layout) onUp() {
	for _, c := range l.getComponents() {
		if f := c.getOnUp(); c.GetActive() && f != nil {
			f()
		}
	}
}

func (l *layout) onDown() {
	for _, c := range l.getComponents() {
		if f := c.getOnDown(); c.GetActive() && f != nil {
			f()
		}
	}
}

func (l *layout) onLeft() {
	for _, c := range l.getComponents() {
		if f := c.getOnLeft(); c.GetActive() && f != nil {
			f()
		}
	}
}

func (l *layout) onRigth() {
	for _, c := range l.getComponents() {
		if f := c.getOnRight(); c.GetActive() && f != nil {
			f()
		}
	}
}
