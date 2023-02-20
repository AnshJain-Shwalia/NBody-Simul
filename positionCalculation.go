package main

import (
	"math"
	"sync"
	"time"
)

// this changeGeneratorCon Uses concurrency(using goroutines).
func changeGeneratorCon(changes chan<- [numObjects][2]float32, bodies [numObjects]body) {
	for {
		newChange := [numObjects][2]float32{}
		acceleration := [numObjects][2]float64{}
		before := time.Now()
		for i := 0; i < len(bodies); i++ {
			b2 := bodies[i]
			index := i
			wg := sync.WaitGroup{}
			for start := 0; start < len(bodies); start += runLength {
				end := start + runLength - 1
				wg.Add(1)
				go accelerate(&acceleration, &bodies, b2, index, start, end, &wg)
			}
			wg.Wait()
		}
		pR(time.Now().Sub(before).Microseconds())
		for j := 0; j < len(bodies); j++ {
			velChangeX := timeStep * acceleration[j][0]
			velChangeY := timeStep * acceleration[j][1]
			bodies[j].velocityVec[0] += velChangeX
			bodies[j].velocityVec[1] += velChangeY
		}

		for k := 0; k < len(bodies); k++ {
			initialPosX := bodies[k].positionVec[0]
			initialPosY := bodies[k].positionVec[1]
			newPosX := initialPosX + (timeStep * bodies[k].velocityVec[0])
			newPosY := initialPosY + (timeStep * bodies[k].velocityVec[1])
			bodies[k].positionVec = [2]float64{newPosX, newPosY}
			newChange[k] = [2]float32{float32(newPosX), float32(newPosY)}
		}
		//pR(len(changes))
		changes <- newChange

	}
}

func changeGeneratorCon2(changes chan<- [numObjects][2]float32, bodies [numObjects]body) {
	for {
		newChange := [numObjects][2]float32{}
		acceleration := [numObjects][2]float64{}
		before := time.Now()
		for i := 0; i < len(bodies); i++ {
			b2 := bodies[i]
			index := i
			done := make(chan bool, numObjects/runLength+1)
			counter := 0
			for start := 0; start < len(bodies); start += runLength {
				end := start + runLength - 1
				counter += 1
				go accelerate2(&acceleration, &bodies, b2, index, start, end, done, numObjects)
			}
			l := 0
			for l < counter {
				<-done
				l += 1
			}
		}
		pR(time.Now().Sub(before).Microseconds())
		for j := 0; j < len(bodies); j++ {
			velChangeX := timeStep * acceleration[j][0]
			velChangeY := timeStep * acceleration[j][1]
			bodies[j].velocityVec[0] += velChangeX
			bodies[j].velocityVec[1] += velChangeY
		}

		for k := 0; k < len(bodies); k++ {
			initialPosX := bodies[k].positionVec[0]
			initialPosY := bodies[k].positionVec[1]
			newPosX := initialPosX + (timeStep * bodies[k].velocityVec[0])
			newPosY := initialPosY + (timeStep * bodies[k].velocityVec[1])
			bodies[k].positionVec = [2]float64{newPosX, newPosY}
			newChange[k] = [2]float32{float32(newPosX), float32(newPosY)}
		}
		//pR(len(changes))
		changes <- newChange

	}
}

func changeGeneratorCon3(changes chan<- [numObjects][2]float32, bodies [numObjects]body) {
	for {
		newChange := [numObjects][2]float32{}
		acceleration := [numObjects][2]float64{}
		before := time.Now()
		for i := 0; i < len(bodies); i++ {
			b2 := bodies[i]
			done := make(chan bool, numObjects/runLength+1)
			counter := 0
			for start := 0; start < len(bodies); start += runLength {
				end := start + runLength
				end = min(end, numObjects)
				counter += 1
				go accelerate3(acceleration[start:start+runLength], bodies[start:start+runLength], b2, done)
			}
			for counter > 0 {
				<-done
				counter -= 1
			}
		}
		pR(time.Now().Sub(before).Microseconds())
		for j := 0; j < len(bodies); j++ {
			velChangeX := timeStep * acceleration[j][0]
			velChangeY := timeStep * acceleration[j][1]
			bodies[j].velocityVec[0] += velChangeX
			bodies[j].velocityVec[1] += velChangeY
		}

		for k := 0; k < len(bodies); k++ {
			initialPosX := bodies[k].positionVec[0]
			initialPosY := bodies[k].positionVec[1]
			newPosX := initialPosX + (timeStep * bodies[k].velocityVec[0])
			newPosY := initialPosY + (timeStep * bodies[k].velocityVec[1])
			bodies[k].positionVec = [2]float64{newPosX, newPosY}
			newChange[k] = [2]float32{float32(newPosX), float32(newPosY)}
		}
		//pR(len(changes))
		changes <- newChange

	}
}

func accelerate(acceleration *[numObjects][2]float64, bodies *[numObjects]body, b2 body, index int, start int, end int, wg *sync.WaitGroup) {
	if numObjects-1 < end {
		end = numObjects - 1
	}
	for ; start <= end; start++ {
		if start == index {
			continue
		}
		b1 := &bodies[start]
		aX, aY := accCalc(b1, &b2)
		acceleration[start][0] += aX
		acceleration[start][1] += aY
	}
	wg.Done()
}

func accelerate2(acceleration *[numObjects][2]float64, bodies *[numObjects]body, b2 body, index int, start int, end int, done chan<- bool, numObjects int) {
	if numObjects-1 < end {
		end = numObjects - 1
	}
	for ; start <= end; start++ {
		if start == index {
			continue
		}
		b1 := &bodies[start]
		aX, aY := accCalc(b1, &b2)
		acceleration[start][0] += aX
		acceleration[start][1] += aY
	}
	done <- true
}

func accelerate3(acceleration [][2]float64, bodies []body, b2 body, done chan<- bool) {
	end := len(acceleration)
	for start := 0; start < end; start++ {
		b1 := &bodies[start]
		aX, aY := accCalc(b1, &b2)
		acceleration[start][0] += aX
		acceleration[start][1] += aY
	}
	done <- true
}

func changeGenerator(changes chan<- [numObjects][2]float32, bodies [numObjects]body) {
	for {
		newChange := [numObjects][2]float32{}
		acceleration := [numObjects][2]float64{}
		before := time.Now()
		for i := 0; i < len(bodies); i++ {
			b2 := bodies[i]
			for j := 0; j < len(bodies); j++ {
				b1 := bodies[j]
				aX, aY := accCalc(&b1, &b2)
				acceleration[j][0] += aX
				acceleration[j][1] += aY
			}
		}
		pR(time.Now().Sub(before).Microseconds())
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
func accCalc(b1 *body, b2 *body) (float64, float64) {
	a := b2.positionVec[0] - b1.positionVec[0]
	b := b2.positionVec[1] - b1.positionVec[1]
	//positions := [numObjects][2]float64{}
	//r1 := (math.Sqrt(b1.mass) / math.Sqrt(22/7)) * areaFactor
	r := math.Sqrt((a * a) + (b * b))
	if r <= minAccDistance {
		return 0.0, 0.0
	}
	force := G * ((b1.mass * b2.mass) / ((r * r) + 0))
	acc := force / b1.mass
	accX := acc * (a / r)
	accY := acc * (b / r)
	return accX, accY
}
