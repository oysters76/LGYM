# L System Implementations 

This project was created to learn more about L-Systems and model them. 
This project aims to implement all the L-System variations found in the following wikipedia article: https://en.wikipedia.org/wiki/L-system

Python was used as it has a nice package called 'turtle' to visualize L-Systems, by mapping the strings to turtle commands. 

PRs are welcome!

## Future Increments to the project: 

1. I would be implementing a front-end GUI for L-Systems so anyone can define "symbols", "axioms", so that they run their own "rules" and generate iterations.
2. I would also be looking into "infering l-systems", specifically "tree" based structures via one-shot or few-shot learning. ( I have a novel concept in mind I'm deciding to work on in this repo, probably in a seperate branch called "inference" )
3. I would be moving the "python" code to a seperate folder inside SRC called "prototype" since this was just the proof of concept.
4. Currently all L-Systems in the wikipedia article have been implemented in this repo.

To run a L-System that draws a "tree" like structure, simply use: 
```python
   from draw import draw_fractal_plant_custom
   draw_fractal_plant_custom({"F":"F[+X]F[-X]+X"},angle=40,f=5) 
```

All of the above new ideas would be implemented using "golang" and "ebitengine" 

1. [golang](https://go.dev/)             : Much faster than Python, suitable for writing this simulation project.
2. [ebitengine](https://ebitengine.org/) : For rendering graphics and GUI. 
