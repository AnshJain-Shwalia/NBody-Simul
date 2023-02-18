package main

import "math"

func changeGenerator(changes chan<- [numObjects][2]float32, bodies [numObjects]body) {
	for {
		newChange := [numObjects][2]float32{}
		acceleration := [numObjects][2]float64{}

		for i := 0; i < len(bodies); i++ {
			b2 := bodies[i]
			for j := 0; j < len(bodies); j++ {
				b1 := bodies[j]
				aX, aY := accCalc(&b1, &b2)
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
func accCalc(b1 *body, b2 *body) (float64, float64) {
	a := b2.positionVec[0] - b1.positionVec[0]
	b := b2.positionVec[1] - b1.positionVec[1]
	//positions := [numObjects][2]float64{}
	//r1 := (math.Sqrt(b1.mass) / math.Sqrt(22/7)) * areaFactor
	r := math.Sqrt((a * a) + (b * b))
	if r <= minAccDistance {
		return 0.0, 0.0
	}
	force := G * ((b1.mass * b2.mass) / (r * r))
	acc := force / b1.mass
	accX := acc * (a / r)
	accY := acc * (b / r)
	return accX, accY

}
