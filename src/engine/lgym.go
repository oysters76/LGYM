package main 

import (
	"fmt"
	"unicode/utf8"
	"bytes"
)

func generate_lsystem(axiom string, 
					  rules map[string]string, 
					  iteration int,
					  verbose bool) string {
	current_generation := axiom 
	var next_generation bytes.Buffer 

	for i:=0; i < iteration; i++ {
		size := utf8.RuneCountInString(current_generation)
		for j:=0;j < size; j++{
			ch := current_generation[j] 
			val, exists := rules[string(ch)]
			if (!exists){
				next_generation.WriteString(string(ch))
				continue; 
			}
			next_generation.WriteString(val) 
		}
		if (verbose){
			fmt.Println(current_generation)
		}
		current_generation = next_generation.String()
		next_generation.Reset()
	}

	return current_generation;
}


func main(){
	//example
	rules := make(map[string]string) 
	rules["F"] = "F+F-F-F+F" 
	axiom := "F" 
	gen := generate_lsystem(axiom, rules, 1, false)
	fmt.Println("gen : ", gen)
}