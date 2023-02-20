package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math"
	"time"
)

func displayInitiator(bodies *[numObjects]body, circles []*canvas.Circle, myContainer *fyne.Container) {
	for i, object := range bodies {
		circle := canvas.NewCircle(myColor)
		mass := object.mass
		//radius := (math.Sqrt(mass) / math.Sqrt(22/7)) * areaFactor
		expression := (mass * 3) / (4 * math.Pi)
		//radius := math.Pow(expression, 0.33) * math.Sqrt(areaFactor)
		radius := (math.Pow(expression, 0.33) * math.Sqrt(areaFactor)) / scaling
		circleDiameter := radius * 2
		circle.Resize(fyne.NewSize(float32(circleDiameter), float32(circleDiameter)))
		circle.StrokeColor = myStrokeColor
		circle.StrokeWidth = 1
		posVec := object.positionVec
		x := posVec[0]
		y := posVec[1]
		circle.Move(fyne.Position{X: float32(x) - float32(radius), Y: float32(y) - float32(radius)})
		myContainer.Add(circle)
		circles[i] = circle
	}
}

func animator(changes <-chan [numObjects][2]float32, circles []*canvas.Circle, counter chan<- bool) {
	for {
		newChange := <-changes
		for i, val := range newChange {
			x := val[0]/float32(scaling) + canvasWidth/2
			y := val[1]/float32(scaling) + canvasHeight/2
			circle := circles[i]
			radius := (circle.Size().Width) / 2
			diameter := radius * 2
			circle.Move(fyne.Position{X: x - radius, Y: y - radius})
			circle.Resize(fyne.NewSize(diameter, diameter))
		}
		for _, circle := range circles {
			circle.Refresh()
		}
		counter <- true
	}
}

func timeStepCounter(counter <-chan bool) {
	beforeTime := time.Now()
	count := 0
	for {
		var _ bool = <-counter
		count += 1
		if count%10 == 0 {
			afterTime := time.Now()
			timeSpan := afterTime.Sub(beforeTime).Nanoseconds()
			var timeSpanSec float64 = float64(timeSpan) / 1e+9
			var Second float64 = 1
			var fps = 10 * (Second / timeSpanSec)
			//pR("fps -> ", fps)
			fps += 1
			timeRate := fps * timeStep
			timeRate += 1
			//pR("timeRate->", timeRate)
			count = 0
			beforeTime = afterTime
		}
	}
}
