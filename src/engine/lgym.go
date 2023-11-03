package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

const CMD_FORWARD = "FORWARD"   //go forward
const CMD_BACKWARD = "BACKWARD" //go backward
const CMD_LEFT = "LEFT"         //turn left at angle 'a'
const CMD_RIGHT = "RIGHT"       //turn right at angle 'a'
const CMD_PUSH = "PUSH"         //push or save state (location/angle)
const CMD_POP = "POP"           //pop or load state (location/angle)
var ALL_COMMANDS = []string{CMD_FORWARD, CMD_BACKWARD, CMD_LEFT, CMD_RIGHT,
	CMD_POP, CMD_PUSH}

func generate_lsystem(lconfig *LConfig,
	verbose bool) string {
	current_generation := lconfig.axiom
	var next_generation bytes.Buffer

	for i := 0; i < int(lconfig.iteration); i++ {
		size := utf8.RuneCountInString(current_generation)
		for j := 0; j < size; j++ {
			ch := current_generation[j]
			val, exists := lconfig.rules[string(ch)]
			if !exists {
				next_generation.WriteString(string(ch))
				continue
			}
			next_generation.WriteString(val)
		}
		if verbose {
			fmt.Println(current_generation)
		}
		current_generation = next_generation.String()
		next_generation.Reset()
	}

	return current_generation
}

func generate_program(lsys string, lconfig *LConfig) []string {
	size := utf8.RuneCountInString(lsys)
	var prog = make([]string, 0, 100)
	for i := 0; i < size; i++ {
		ch := lsys[i]
		val, exists := lconfig.cmds[string(ch)]
		if !exists {
			continue
		}
		prog = append(prog, val)
	}
	return prog
}

func RunLSystem(filename string) {
	lconfig := Parse(filename)
	gen := generate_lsystem(&lconfig, false) // generate l-system string
	prog := generate_program(gen, &lconfig)  // convert symbols to graphic program commands
	Renderlsystem(prog, len(prog), &lconfig) // render graphics via commands
}

func main() {
	RunLSystem("examples/fractal.json")
}
