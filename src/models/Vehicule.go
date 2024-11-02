package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Vehicule struct {
	posX, posY float32
	angle int
	status bool
	parked bool
	entranceChannel chan bool
	observers []Observer
}

func NewVehicule() *Vehicule {
	return &Vehicule{posX: 0, posY: 0, angle: 270, parked: false, status: true, entranceChannel: make(chan bool)}
}

func wait() {
	time.Sleep(2 * time.Millisecond)
}

func (v *Vehicule) Run (widthContainer, heightContainer float32, entrance *Entrance, parkingSlots *[20]ParkingSlot){

	v.posX = widthContainer * 0.80 
	v.posY = 0

	for v.status {

		if !v.parked{
			
			if entrance.posY == v.posY + 20 {
				entrance.AddToIncomingQueue(*v)
	
				select {
				case <- v.entranceChannel:
					v.angle = 180
					parkingSlot := getAvailableParkingSlot(parkingSlots)
					for !v.parked {

						v.posX -= 1
						v.NotifyAll()
						if v.posX == entrance.posX - 100 {
							fmt.Println("ya pase la entrada")
							entrance.NotifyVehiculePassed()
							fmt.Println("comprobacion de no deadlock")
						}
						if v.posX == parkingSlot.posX {
							
							for !v.parked {
								if parkingSlot.posY > v.posY {
									v.posY += 1
									v.angle = 270
								} else if parkingSlot.posY < v.posY {
									v.posY -= 1
									v.angle = 90
								} else if parkingSlot.posY == v.posY {
									v.parked = true
									fmt.Print("ya me estacione")
									time.Sleep( time.Duration((rand.Intn(2) + 3) * int(time.Second)))
									parkingSlot.busy = false
								}
								wait()
								v.NotifyAll()
							}
						}
						wait()
					}
				}
				
			} else {
				v.posY += 1
			}
			v.NotifyAll()
			wait()
		} else {

			for v.posY != entrance.posY + 20 {
				if v.posY > entrance.posY + 20 {
					v.angle = 90
					v.posY -= 1
					v.NotifyAll()
				} else if v.posY < entrance.posY + 20 {
					v.angle = 270
					v.posY += 1
					v.NotifyAll()
				}
				wait()
			}
			fmt.Println("a la altura de la entrada")
			v.angle = 0
			for v.posX != entrance.posX - 100 {
				v.posX += 1
				v.NotifyAll()
				wait()
			}
			entrance.AddToOutgoingQueue(*v)

			select {
			case <- v.entranceChannel:
				for v.posX != widthContainer *.80 {
					v.posX += 1
					v.NotifyAll()
					wait()
				}
				entrance.NotifyVehiculePassed()
				v.angle = 270
				for v.status {
					v.posY += 1
					v.NotifyAll()
					wait()
					if v.posY == heightContainer {
						v.status = false
					}
				}
			}
		}
	}
}

func (v *Vehicule) Register (observer Observer){
	v.observers = append(v.observers, observer)
}

func (v *Vehicule) Unregister (observer Observer){
	for i, o := range v.observers {
		if o == observer{
			v.observers = append(v.observers[:i], v.observers[i+1:]... )
		}
	}
}

func (v *Vehicule) SetStatus(status bool){
	v.status = status
}

func (v *Vehicule) NotifyAll() {
	for _, observer := range v.observers{
		observer.Update(v.posX, v.posY, v.angle)
	}
}

func getAvailableParkingSlot (parkingSlots *[20]ParkingSlot) *ParkingSlot {

	// var founded *ParkingSlot
	
	// for _, parkingSlot := range parkingSlots {
	// 	if !parkingSlot.busy {
	// 		parkingSlot.busy = true
	// 		founded = parkingSlot
	// 	}
	// }

	// return founded
	for i := 0; i < len(parkingSlots); i++ {
        if !parkingSlots[i].busy {
            parkingSlots[i].busy = true
            return &parkingSlots[i] // Retornamos la referencia al espacio de estacionamiento
        }
    }
	return nil
	// return parkingSlots[0]
}