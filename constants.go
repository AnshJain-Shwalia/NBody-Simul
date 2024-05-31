package main

import "image/color"

const timeStep float64 = 0.0001
const G float64 = 50
const canvasHeight = 800
const canvasWidth = 1200
const numObjects int = 20
const centerX float64 = 0 //float64(canvasWidth) / 2
const centerY float64 = 0 //float64(canvasHeight) / 2
const minDistance float64 = 100
const maxDistance float64 = 150
const minVelocity float64 = 10
const maxVelocity float64 = 15
const minMass float64 = 10
const maxMass float64 = 15
const chanSize int = 100
const areaFactor float64 = 8
const scaling float64 = 3
const minAccDistance float64 = 5
const runLength int = 50

var myColor color.NRGBA = color.NRGBA{R: 255, G: 255, B: 255, A: 15}
var myStrokeColor color.NRGBA = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

type body struct {
	mass        float64
	positionVec [2]float64
	velocityVec [2]float64
}

var Sun body = body{mass: 10000, velocityVec: [2]float64{0.0, 0.0}, positionVec: [2]float64{centerX, centerY}}
