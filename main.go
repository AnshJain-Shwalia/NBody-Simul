package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	//bodies contains the objects which are to be operated upon
	bodies := random()
	//bodies[len(bodies)-1] = Sun
	//pR(bodies)
	myApp := app.New()
	myWindow := myApp.NewWindow("Simulation")
	myContainer := container.NewWithoutLayout()
	// circles holds the references to the circles which will later need to be updated
	circles := make([]*canvas.Circle, numObjects)
	displayInitiator(&bodies, circles, myContainer)
	myWindow.Resize(fyne.NewSize(float32(canvasWidth), float32(canvasHeight)))
	myWindow.SetContent(myContainer)
	changes := make(chan [numObjects][2]float32, chanSize)
	go changeGenerator(changes, bodies)
	counter := make(chan bool, 10)
	go animator(changes, circles, counter)
	go timeStepCounter(counter)
	myWindow.Show()
	myApp.Run()
}
