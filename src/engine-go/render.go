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

/*
Structure that stores the position and angle at current timestep in PUSH.
POP retrieves this state.
*/
type State struct {
	x       int32
	y       int32
	c_angle int32
}

type Vec2d struct {
	x float64
	y float64
}

func Rotate(vec Vec2d, angle int32) Vec2d {
	//     perp = Vec2D(-self[1], self[0])
	//     angle = math.radians(angle)
	//     c, s = math.cos(angle), math.sin(angle)
	//     return Vec2D(self[0]*c+perp[0]*s, self[1]*c+perp[1]*s)
	perp := Vec2d{-vec.y, vec.x}
	rad := ToRad(angle)
	c, s := math.Sincos(rad)
	return Vec2d{vec.x*c + perp.x*s, vec.y*c + perp.y*s}
}

/*
This function renders a l-system given the program commands.
prog     : the array of commands (e.g. FORWARD, BACKWARD, LEFT, RIGHT, POP, PUSH)
size     : the size of the array of commands
lconfig  : the L-System configuration
*/
func Renderlsystem(prog []string, size int,
	lconfig *LConfig) {
	rl.InitWindow(lconfig.width, lconfig.height, lconfig.title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	var x int32 = lconfig.ix
	var y int32 = lconfig.iy
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
			dx, dy = PolarToXY(-lconfig.magintude, ToRad(c_angle))
		case CMD_FORWARD:
			shouldDrawLine = true
			dx, dy = PolarToXY(lconfig.magintude, ToRad(c_angle))
		case CMD_RIGHT:
			c_angle = -lconfig.angle
		case CMD_LEFT:
			c_angle = lconfig.angle
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

func TestDrawing() {
	rl.InitWindow(1200, 800, "helloworld")
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		dx, dy := PolarToXY(50, ToRad(360))
		DrawLine(1200/2, 800/2, 1200/2+dx, 800/2+dy)
		rl.EndDrawing()
	}
}
