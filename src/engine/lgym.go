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

func generate_lsystem(axiom string,
	rules map[string]string,
	iteration int,
	verbose bool) string {
	current_generation := axiom
	var next_generation bytes.Buffer

	for i := 0; i < iteration; i++ {
		size := utf8.RuneCountInString(current_generation)
		for j := 0; j < size; j++ {
			ch := current_generation[j]
			val, exists := rules[string(ch)]
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

func generate_program(lsys string, command_map map[string]string) []string {
	size := utf8.RuneCountInString(lsys)
	var prog = make([]string, 0, 100)
	for i := 0; i < size; i++ {
		ch := lsys[i]
		val, exists := command_map[string(ch)]
		if !exists {
			fmt.Println("[ERROR] Incompatible L-System with given command map!")
			break
		}
		prog = append(prog, val)
	}
	return prog
}

func main() {
	//example for generating an l-system and generating it's draw program.
	rules := make(map[string]string)
	rules["F"] = "F+F-F-F+F"
	axiom := "F"
	gen := generate_lsystem(axiom, rules, 10, false)

	prog_rules := make(map[string]string)
	prog_rules["F"] = CMD_FORWARD
	prog_rules["+"] = CMD_LEFT
	prog_rules["-"] = CMD_RIGHT
	prog := generate_program(gen, prog_rules)

	Renderlsystem(prog, len(prog), 5, 1500, 1000, 60, 500, 500)
}
