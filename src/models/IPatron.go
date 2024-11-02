package models

type Observer interface {
	Update(x, y float32, angle int)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

