package views

import (
	// "fmt"
	"math/rand"
	"simulador/src/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)
 
type Vehicule struct {
	Vehicule *canvas.Image
}

func NewVehicule () *Vehicule{
	return &Vehicule{Vehicule: nil}
}

func randomVehiculeImage () string{
	var options [3]string
	options[0] = "auto1.png"
	options[1] = "auto2.png"
	options[2] = "auto3.png"
	i := rand.Intn(2)

	return options[i]
}

func (v *Vehicule) AddVehicule(sc *scenes.Scene, widthContainer float32) {
	vehicule := canvas.NewImageFromURI(storage.NewFileURI("./src/assets/autoDown.png"))
	vehicule.Resize(fyne.NewSize(100, 100))
	vehicule.Move(fyne.NewPos(widthContainer * .875, 0))

	v.Vehicule = vehicule
	sc.AddImage(vehicule)
}

func (v *Vehicule) Update(x, y float32, angle int) {
	// fmt.Printf("%f : %f\n", x, y)
	v.Vehicule.Move(fyne.NewPos(x, y))
	switch angle {
	case 90:
		v.Vehicule.File = "./src/assets/autoUp.png"
	case 180:
		v.Vehicule.File = "./src/assets/autoLeft.png"
	case 270:
		v.Vehicule.File = "./src/assets/autoDown.png"
	case 0: 
		v.Vehicule.File = "./src/assets/autoRight.png"
	}
	v.Vehicule.Refresh()
	
}