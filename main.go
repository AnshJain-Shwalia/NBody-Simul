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
	//b1 := body{mass: 10000, positionVec: [2]float64{centerX, centerY}, velocityVec: [2]float64{float64(0), float64(0)}}
	//posVector, velVector := randomPerpendicularUnitVectors()
	//posVector[0] = (posVector[0] * 100) + centerX
	//posVector[1] = (posVector[1] * 100) + centerY
	//velVector[0] *= 10
	//velVector[1] *= 10
	//b2 := body{mass: 100, positionVec: posVector, velocityVec: velVector}
	//bodies = [2]body{b1, b2}
	pR(bodies)
	myApp := app.New()
	myWindow := myApp.NewWindow("Circles")
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
