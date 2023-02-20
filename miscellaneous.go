package main

import (
	"fmt"
	"math"
	"math/rand"
)

//func intersectionCheck(b1 body, b2 body) bool {
//
//}

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
	//velVector[0] = velVector[0] * -1
	if rand.Intn(2) == 0 {
		velVector[0] = velVector[0] * -1
	} else {
		velVector[1] = velVector[1] * -1
	}
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
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
