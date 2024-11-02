package models

import "fmt"

type Entrance struct {
	posX, posY float32
	busy bool
	isIncomig bool
	isOutcoming bool
	incomingVehiculeChan chan Vehicule
	outgoingVehiculeChan chan Vehicule
	incomingQueue, outgoingQueue []Vehicule
	hasPassed chan bool
}

func NewEntrace (x, y float32) *Entrance{
	return &Entrance{
		posX: x, 
		posY: y, 
		busy: false, 
		isIncomig: false, 
		isOutcoming: false, 
		incomingQueue: []Vehicule{}, 
		outgoingQueue:  []Vehicule{},
		incomingVehiculeChan: make(chan Vehicule),
		outgoingVehiculeChan: make(chan Vehicule),
		hasPassed: make(chan bool),
	}
}

func (e *Entrance) NotifyVehiculePassed () {
	e.busy = false
	e.hasPassed <- true
}

func (e *Entrance) GetFirstIncomingVehicule () *Vehicule{
	if(len(e.incomingQueue) >= 1){
		first := e.incomingQueue[0]
		e.incomingQueue = e.incomingQueue[1:]
		return &first
	} else {
		return nil
	}
}

func (e *Entrance) GetFirstOutgoingVehicule () *Vehicule{
	if len(e.outgoingQueue) >= 1 {
		first := e.outgoingQueue[0]
		e.outgoingQueue = e.outgoingQueue[1:]
		return &first
	} else {
		return nil
	}
}

func (e *Entrance) AddToIncomingQueue (vehicule Vehicule){
	e.incomingVehiculeChan <- vehicule
}

func (e *Entrance) AddToOutgoingQueue (vehicule Vehicule){
	e.outgoingVehiculeChan <- vehicule
}

func (e *Entrance) Release () {
	e.busy = false
	e.isIncomig = false
	e.isOutcoming = false
}

func (e *Entrance) LetNextIncomingVehicule () {
	next := e.GetFirstIncomingVehicule()
	if next != nil {
		e.busy = true
		e.isIncomig = true
		next.entranceChannel <- true
	} else if e.isIncomig {
		fmt.Println("ya no tengo mas entrando")
		e.isIncomig = false
		e.LetNextOutgoingVehicule()
	} else {
		fmt.Print("dejando libre desde entrante")
		e.Release()
	}
}

func (e *Entrance) LetNextOutgoingVehicule () {
	next := e.GetFirstOutgoingVehicule()
	if next != nil  {
		e.busy = true
		e.isOutcoming = true
		next.entranceChannel <- true
	} else if e.isOutcoming {
		fmt.Println("ya no tengo mas saliendo")
		e.isOutcoming = false
		e.LetNextIncomingVehicule()
	} else {
		fmt.Print("dejando libre desde saliente")
		e.Release()
	}
}

func (e *Entrance) Run (){

	for {
		select {
	
			case passed := <- e.hasPassed:
				fmt.Println("ya salio uno", passed)
				if e.isIncomig {
					e.LetNextIncomingVehicule()
				} else if e.isOutcoming{
					e.LetNextOutgoingVehicule()
				}
			
			case vehicule := <- e.incomingVehiculeChan:
				fmt.Println("ya entro uno")
				e.incomingQueue = append(e.incomingQueue, vehicule)
				if(!e.isIncomig || e.isIncomig) && !e.isOutcoming && !e.busy {
					if !e.isIncomig {
						e.busy = true
						e.isIncomig = true
					}
					first := e.GetFirstIncomingVehicule()
					first.entranceChannel <- true
				}
			
			case vehicule := <- e.outgoingVehiculeChan:
				fmt.Println("entro uno para salir")
				e.outgoingQueue = append(e.outgoingQueue, vehicule)
				if(!e.isOutcoming || e.isOutcoming) && !e.isIncomig && !e.busy {
					if !e.isOutcoming {
						e.busy = true
						e.isOutcoming = true
					}
	
					first := e.GetFirstOutgoingVehicule()
					first.entranceChannel <- true
				}
		}
	}

}