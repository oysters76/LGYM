# LGYM

This project was created to learn more about L-Systems and model them. 
This project aims to implement all the L-System variations found in the following wikipedia article: https://en.wikipedia.org/wiki/L-system

Python was used as it has a nice package called 'turtle' to visualize L-Systems, by mapping the strings to turtle commands. 

PRs are welcome!

## Future Increments to the project: 

1. I would be implementing a front-end GUI for L-Systems so anyone can define "symbols", "axioms", so that users run their own "rules" and generate iterations based on those rules using "raylib" 💭 
2. I would also be looking into "infering l-systems", specifically "tree" based structures via one-shot or few-shot learning. ( I have a novel concept in mind I'm deciding to work on in this repo, probably in a seperate branch called "inference" ) 💭
3. I would be moving the "python" code to a seperate folder inside SRC called "prototype" since this was just the proof of concept. ✅
4. Currently all L-Systems in the wikipedia article have been implemented in this repo. ✅
5. The L-System generator is now written using GO (ported from Python) ✅

To run a L-System that draws a "tree" like structure, simply use: 
```python
   from draw import draw_fractal_plant_custom
   draw_fractal_plant_custom({"F":"F[+X]F[-X]+X"},angle=40,f=5) 
```

All of the above new ideas would be implemented using "golang" and "raylib" (go has a raylib binding)

1. [golang](https://go.dev/)             : Much faster than Python, suitable for writing this simulation project.
2. [raylib](https://github.com/gen2brain/raylib-go) : For rendering graphics and GUI. 


## Basic Usage 
The following simple golang program illustrates the process in which this tool can be used: 

```go
func main(){
   lconfig := Parse("examples/fractal.json") //load l-system config
   gen := generate_lsystem(&lconfig, false)  // generate l-system string
   prog := generate_program(gen, &lconfig)   // convert symbols to graphic program commands
   Renderlsystem(prog, len(prog), &lconfig)  // render graphics via commands
}
```