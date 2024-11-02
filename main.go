package main

import (
	"fmt"
	"simulador/src/models"
	"simulador/src/scenes"
	"simulador/src/views"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"math/rand"
)

func main(){
	
	a := app.New()
    w := a.NewWindow("Simulador")
	width := 1280
	height := 720
    w.Resize(fyne.NewSize(float32(width), float32(height))) // Tamaño total de la ventana

	scene := scenes.NewScene(w)
	parkingSlots, entrance := scene.Init(width, height)
	fmt.Print(parkingSlots, entrance)
	
	go func () {
		i:= 1
		loop := true
		go entrance.Run()

		for loop {
			
			v1 := models.NewVehicule()
			vehicule := views.NewVehicule()
			vehicule.AddVehicule(scene, float32(width))
			v1.Register(vehicule)
	
			go v1.Run(float32(width), float32(height), entrance, parkingSlots)
	
			i++
			time.Sleep( time.Duration(rand.Intn(1000) + 1000) * time.Millisecond)
			fmt.Println(i, "hola")
			if i == 100 {
				loop = false
			}
		}
	}()

	w.ShowAndRun()
}