package core

//Style ключ стиля
type Style byte

const (
	Border = iota
	BackgroundColor
	TextColor
	Wigth
	Heigth
	ZeroX
	ZeroY
)

//IComponent итерфейс компонента
type IComponent interface {
	OnCreate(f func()) IComponent
	OnUpdate(f func()) IComponent
	OnDelete(f func()) IComponent
	OnSelect(f func()) IComponent
	OnActive(f func()) IComponent
	OnBack(f func()) IComponent
	SetComponents(comp ...IComponent) IComponent
	SetStyle(style map[Style]interface{}) IComponent
	GetStyle() map[Style]interface{}
	SetActive(active bool) IComponent
	GetActive() bool
	getComponents() []IComponent
	getCreate() func()
	getDelete() func()
	getOnEnter() func()
	getOnTab() func()
	getOnBackSpace() func()
	getOnUp() func()
	getOnDown() func()
	getOnLeft() func()
	getOnRight() func()
}

type component struct {
	active      bool
	components  []IComponent
	style       map[Style]interface{}
	onCreate    func()
	onUpdate    func()
	onDelete    func()
	onTab       func()
	onEnter     func()
	onBackSpace func()
	onUp        func()
	onDown      func()
	onLeft      func()
	onRigth     func()
}

//Dev создание блока
func Dev() IComponent {
	comp := &component{}
	return comp
}

func (c *component) OnCreate(f func()) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.onCreate = f
	return c
}

func (c *component) OnUpdate(f func()) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.onUpdate = f
	return c
}

func (c *component) OnDelete(f func()) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.onDelete = f
	return c
}

func (c *component) OnSelect(f func()) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.onTab = f
	return c
}

func (c *component) OnActive(f func()) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.onEnter = f
	return c
}

func (c *component) OnBack(f func()) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.onBackSpace = f
	return c
}
func (c *component) SetComponents(comp ...IComponent) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.components = comp
	if f := c.onUpdate; f != nil {
		f()
	}
	return c
}

func (c *component) SetStyle(style map[Style]interface{}) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.style = style
	if f := c.onUpdate; f != nil {
		f()
	}
	return c
}

func (c *component) GetStyle() map[Style]interface{} {
	rw.Lock()
	defer rw.Unlock()

	return c.style
}

func (c *component) SetActive(active bool) IComponent {
	rw.Lock()
	defer rw.Unlock()

	c.active = active
	return c
}

func (c *component) GetActive() bool {
	rw.Lock()
	defer rw.Unlock()

	return c.active
}

func (c *component) getComponents() []IComponent {
	rw.Lock()
	defer rw.Unlock()

	return c.components
}

func (c *component) getCreate() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onCreate
}

func (c *component) getDelete() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onDelete
}

func (c *component) getOnEnter() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onEnter
}

func (c *component) getOnTab() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onTab
}

func (c *component) getOnBackSpace() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onBackSpace
}

func (c *component) getOnUp() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onUp
}

func (c *component) getOnDown() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onDown
}

func (c *component) getOnLeft() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onLeft
}

func (c *component) getOnRight() func() {
	rw.Lock()
	defer rw.Unlock()

	return c.onRigth
}
