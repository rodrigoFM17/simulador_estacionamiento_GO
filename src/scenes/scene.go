package scenes

import (
	"fmt"
	"image/color"
    "simulador/src/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)


type Scene struct {
	scene fyne.Window
	container *fyne.Container 
}

func NewScene(scene fyne.Window) *Scene {
	return &Scene{scene: scene, container: nil}
}

func addLinesStreet (container *fyne.Container, widthContainer float32) {

    width := 10
    height := 50
    for i := 0; i*height*2 < 720 ; i++ {
        line := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        line.Resize(fyne.NewSize(float32(width), float32(height)))
        fmt.Print(i * height + 50)
        line.Move(fyne.NewPos(widthContainer * .875 - float32(width) / 2, float32(i * height *2 + 50)))
        container.Add(line)
    }
}

func addParkingSlots (container *fyne.Container, widthContainer, heightContainer float32) [20]models.ParkingSlot{
    var padding float32 = 50
    width := widthContainer * .75 - padding * 2
    var slotThickness float32 = 5
    slotWidth := width /10
    var slotHeight float32 = 100

    var parkingSlots [20]models.ParkingSlot

    for  i := 0; i <10; i++ {

        left := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        left.Resize(fyne.NewSize(slotThickness, slotHeight))
        left.Move(fyne.NewPos(padding + float32(i) * slotWidth, padding))
        top := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        top.Resize(fyne.NewSize(slotWidth, slotThickness))
        top.Move(fyne.NewPos(padding + float32(i) * slotWidth, padding))
        right := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        right.Resize(fyne.NewSize(slotThickness, slotHeight))
        right.Move(fyne.NewPos(padding + float32(i+1) * slotWidth, padding))
        container.Add(left)
        container.Add(top)
        container.Add(right)
        parkingSlots[i] = *models.NewParkingSlot(
            padding + float32(i) * slotWidth + slotThickness,
            padding,
            slotWidth - slotThickness * 2,
            slotHeight - slotThickness,
        )
    }

    for  i := 0; i <10; i++ {

        left := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        left.Resize(fyne.NewSize(slotThickness, slotHeight))
        left.Move(fyne.NewPos(padding + float32(i) * slotWidth, heightContainer - padding - slotHeight))
        top := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        top.Resize(fyne.NewSize(slotWidth, slotThickness))
        top.Move(fyne.NewPos(padding + float32(i) * slotWidth, heightContainer - padding))
        right := canvas.NewRectangle(color.RGBA{236, 240, 7, 255})
        right.Resize(fyne.NewSize(slotThickness, slotHeight))
        right.Move(fyne.NewPos(padding + float32(i+1) * slotWidth, heightContainer - padding - slotHeight))
        container.Add(left)
        container.Add(top)
        container.Add(right)
        parkingSlots[i + 10] = *models.NewParkingSlot(
            padding + float32(i) * slotWidth + slotThickness,
            heightContainer - padding - 100,
            slotWidth - slotThickness * 2,
            slotHeight - slotThickness,
        )
    }

    fmt.Print(parkingSlots)
    return parkingSlots
}

func (s *Scene) Init(width, height int) (*[20]models.ParkingSlot, *models.Entrance){

	s.container = container.NewWithoutLayout()

    leftBackground := canvas.NewRectangle(color.RGBA{102, 102, 102, 255})
    leftBackground.Resize(fyne.NewSize(float32(width) * .75, float32(height))) 
    leftBackground.Move(fyne.NewPos(0, 0))

    wall := canvas.NewRectangle(color.RGBA{171, 58, 50, 255})
    wall.Resize(fyne.NewSize(20, float32(height)))
    wall.Move(fyne.NewPos(float32(width) * .75 - 20, 0))

    rightBackground := canvas.NewRectangle(color.RGBA{39, 143, 44, 255})
    rightBackground.Resize(fyne.NewSize(float32(width) * 0.25, float32(height))) 
    rightBackground.Move(fyne.NewPos(float32(width) * .75, 0))   

    street := canvas.NewRectangle(color.RGBA{102, 102, 102, 255})
    street.Resize(fyne.NewSize(float32(width) * 0.15 , float32(height)))
    street.Move(fyne.NewPos(float32(width) * 0.80, 0))

    entranceToParking := canvas.NewRectangle(color.RGBA{102, 102, 102, 255})
    entranceToParking.Resize(fyne.NewSize(float32(width) * 0.15 + 20, 100))
    entranceToParking.Move(fyne.NewPos(float32(width) *.75 - 20, float32(height) / 2 - 50))
    entrance := models.NewEntrace(float32(width)*.75, float32(height) / 2 - 50)

    s.container.Add(leftBackground)
    s.container.Add(rightBackground)
    s.container.Add(street)
    s.container.Add(wall)
    s.container.Add(entranceToParking)
    addLinesStreet(s.container, float32(width))
    parkingSlots := addParkingSlots(s.container, float32(width), float32(height))

    s.scene.SetContent(s.container)
    return &parkingSlots, entrance
}

func (s *Scene) AddImage (image *canvas.Image){
    s.container.Add(image)
    s.container.Refresh()
}