/*
Render engine for LGYM
----------------------

Creates a window that parses a given JSON formatted instruction file and then it
draws the generated L-System on the Window.

The JSON document needs to have information about the rules, axiom, and the mappings
between the symboles/constants to commands, and the iteration #.
*/
package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func PolarToXY(mag int32, angle float64) (int32, int32) {
	sin_v, cos_v := math.Sincos(angle)
	return int32(float64(mag) * cos_v), int32(float64(mag) * sin_v)
}

func ToRad(angle int32) float64 {
	return float64(angle) * (math.Pi / 180)
}

func DrawLine(x int32, y int32, x2 int32, y2 int32) {
	rl.DrawLine(x, y, x2, y2, rl.RayWhite)
}

type State struct {
	x       int32
	y       int32
	c_angle int32
}

func Renderlsystem(prog []string, size int, magnitude int32, width int32, height int32, angle int32, ix int32, iy int32) {
	rl.InitWindow(width, height, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	var x int32 = ix
	var y int32 = iy
	var dx int32 = 0
	var dy int32 = 0
	var c_angle int32 = 0
	var prog_stack []State
	ind := 0
	for !rl.WindowShouldClose() {
		if ind >= size {
			continue
		}
		rl.BeginDrawing()
		shouldDrawLine := false
		switch prog[ind] {
		case CMD_BACKWARD:
			shouldDrawLine = true
			dx, dy = PolarToXY(-magnitude, ToRad(c_angle))
		case CMD_FORWARD:
			shouldDrawLine = true
			dx, dy = PolarToXY(magnitude, ToRad(c_angle))
		case CMD_RIGHT:
			c_angle += -angle
		case CMD_LEFT:
			c_angle += angle
		case CMD_PUSH:
			prog_stack = append(prog_stack, State{x: x, y: y, c_angle: c_angle})
		case CMD_POP:
			l := len(prog_stack)
			state := prog_stack[l-1]
			prog_stack = prog_stack[:l-1]
			x = state.x
			y = state.y
			c_angle = state.c_angle
		}
		if shouldDrawLine {
			DrawLine(x, y, x+dx, y+dy)
			x += dx
			y += dy
		}
		ind += 1
		rl.EndDrawing()
	}
}
