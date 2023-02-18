package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const timeStep float64 = 0.001
const G float64 = 50
const canvasHeight = 800
const canvasWidth = 1200
const numObjects int = 100
const centerX float64 = float64(canvasWidth) / 2
const centerY float64 = float64(canvasHeight) / 2
const minDistance float64 = 100
const maxDistance float64 = 150
const minVelocity float64 = 10
const maxVelocity float64 = 15
const minMass float64 = 10
const maxMass float64 = 15
const chanSize int = 100
const areaFactor float64 = 2

var myColor color.NRGBA = color.NRGBA{R: 255, G: 255, B: 255, A: 15}
var myStrokeColor color.NRGBA = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

type body struct {
	mass        float64
	positionVec [2]float64
	velocityVec [2]float64
}

var Sun body = body{mass: 10000, velocityVec: [2]float64{0.0, 0.0}, positionVec: [2]float64{centerX, centerY}}

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
	for i, object := range bodies {
		circle := canvas.NewCircle(myColor)
		mass := object.mass
		//radius := (math.Sqrt(mass) / math.Sqrt(22/7)) * areaFactor
		expression := (mass * 3) / (4 * math.Pi)
		radius := math.Pow(expression, 0.33) * areaFactor
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

func changeGenerator(changes chan<- [numObjects][2]float32, bodies [numObjects]body) {
	for {
		newChange := [numObjects][2]float32{}
		acceleration := [numObjects][2]float64{}
		//velocities := [numObjects][2]float64{}
		//positions := [numObjects][2]float64{}
		for i := 0; i < len(bodies); i++ {
			b2 := bodies[i]
			for j := 0; j < len(bodies); j++ {
				b1 := bodies[j]
				aX, aY := accCalc(b1, b2)
				acceleration[j][0] += aX
				acceleration[j][1] += aY
			}
		}
		for i := 0; i < len(bodies); i++ {
			velChangeX := timeStep * acceleration[i][0]
			velChangeY := timeStep * acceleration[i][1]
			bodies[i].velocityVec[0] += velChangeX
			bodies[i].velocityVec[1] += velChangeY
		}

		for i := 0; i < len(bodies); i++ {
			initialPosX := bodies[i].positionVec[0]
			initialPosY := bodies[i].positionVec[1]
			newPosX := initialPosX + (timeStep * bodies[i].velocityVec[0])
			newPosY := initialPosY + (timeStep * bodies[i].velocityVec[1])
			bodies[i].positionVec = [2]float64{newPosX, newPosY}
			newChange[i] = [2]float32{float32(newPosX), float32(newPosY)}
		}
		//pR(len(changes))
		changes <- newChange
	}
}

// this calculates the acceleration caused on b1 by b2
func accCalc(b1 body, b2 body) (float64, float64) {
	a := b2.positionVec[0] - b1.positionVec[0]
	b := b2.positionVec[1] - b1.positionVec[1]
	//r1 := (math.Sqrt(b1.mass) / math.Sqrt(22/7)) * areaFactor
	r := math.Sqrt((a * a) + (b * b))
	if r <= 5 {
		return 0.0, 0.0
	}
	force := G * ((b1.mass * b2.mass) / (r * r))
	acc := force / b1.mass
	accX := acc * (a / r)
	accY := acc * (b / r)
	return accX, accY

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
			pR("fps -> ", fps)
			timeRate := fps * timeStep
			pR("timeRate->", timeRate)
			count = 0
			beforeTime = afterTime
		}
	}
}

//func intersectionCheck(b1 body, b2 body) bool {
//
//}

func animator(changes <-chan [numObjects][2]float32, circles []*canvas.Circle, counter chan<- bool) {
	for {
		//time.Sleep(time.Millisecond * 1)
		newChange := <-changes
		for i, val := range newChange {
			x := val[0]
			y := val[1]
			circle := circles[i]
			radius := (circle.Size().Width) / 2
			circle.Move(fyne.Position{X: x - radius, Y: y - radius})
		}
		for _, circle := range circles {
			circle.Refresh()
		}
		counter <- true
	}
}

// Returns randomised array of object with random velocities and masses within a certain limit
func random() [numObjects]body {
	returnArray := [numObjects]body{}
	for i := 0; i < numObjects; i++ {
		newBody := body{}
		newBody.mass = randomFloat64(minMass, maxMass, 3)
		posVector, velVector := randomPerpendicularUnitVectors()
		distance := randomFloat64(minDistance, maxDistance, 3)
		velocity := randomFloat64(minVelocity, maxVelocity, 3)
		newBody.positionVec = [2]float64{(posVector[0] * distance) + centerX, (posVector[1] * distance) + centerY}
		newBody.velocityVec = [2]float64{velVector[0] * velocity, velVector[1] * velocity}
		returnArray[i] = newBody
	}
	return returnArray
}

func randomPerpendicularUnitVectors() ([2]float64, [2]float64) {
	posX := randomFloat64(-1.0, 1.0, 3)
	posY := math.Sqrt(1 - (posX * posX))
	if rand.Intn(2) == 0 {
		posY *= -1
	}
	posVector := [2]float64{posX, posY}
	velVector := [2]float64{posVector[1], posVector[0]}
	velVector[0] = velVector[0] * -1
	//if rand.Intn(2) == 0 {
	//	velVector[0] = velVector[0] * -1
	//} else {
	//	velVector[1] = velVector[1] * -1
	//}
	//perpendicularCheck(posVector, velVector)
	return posVector, velVector

}

//
//func perpendicularCheck(vec1 [2]float64, vec2 [2]float64) {
//	result := (vec1[0] * vec2[0]) + (vec1[1] * vec2[1])
//	if !(result < 0.00001 && result > -0.00001) {
//		pR("Not perpendicular", "resul->", result)
//		panic("not perpendicular")
//	}
//}

// Return a random float between [a,b] with some precision
func randomFloat64(a float64, b float64, precision int) float64 {
	offset := a
	if a > b {
		panic("Wrong request value.")
	}
	a *= math.Pow10(precision)
	b *= math.Pow10(precision)
	result := rand.Intn(int(b) - int(a) + 1)
	answer := float64(result)/math.Pow10(precision) + offset
	return answer

}

func pR(val ...interface{}) {
	for i, value := range val {
		if i < len(val)-1 {
			fmt.Printf("%v, ", value)
		} else {
			fmt.Printf("%v\n", value)
		}
	}
}
