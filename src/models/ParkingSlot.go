package models

type ParkingSlot struct {
	posX float32
	posY float32
	width float32
	height float32
	busy bool
}

func NewParkingSlot (x, y, width, height float32) *ParkingSlot {
	return &ParkingSlot{
		posX: x, 
		posY: y, 
		width: width, 
		height: height, 
		busy: false,
	}
}

func (p *ParkingSlot) Ocupy () {
	p.busy = true
}

func (p *ParkingSlot) Leave () {
	p.busy = false
}