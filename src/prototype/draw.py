
from turtle import * 
from lsystem import gen_lsystem 

def draw_dragon(generation):
    clearscreen() 
    mode("logo") 
    speed("fastest")
    up()
    setpos(-200,-200)
    down()
    for token in generation:
        if token == "F" or token == "G":
            forward(15) 
        elif token == "+":
            left(90)
        elif token == "-":
            right(90)
    done() 

def draw_koch(generation):
    clearscreen() 
    speed("fastest")
    up()
    setpos(-200,-200)
    down()
    for token in generation:
        if token == "F":
            forward(2) 
        elif token == "+":
            left(90)
        elif token == "−":
            right(90)
    done()

def draw_si_traingle(generation):
    clearscreen() 
    speed("fastest")
    up()
    setpos(-200,-200)
    down()
    for token in generation:
        if token == "A" or token == "B":
            forward(2) 
        elif token == "+":
            left(60)
        elif token == "−":
            right(60)
    done()

def draw_tree(generation,angle=45):
    clearscreen() 
    speed("fastest")
    mode("logo")
    up()
    setpos(0,-400)
    down()
    stack = [] 
    for token in generation:
        if token == "0":
            forward(1) 
        if token == "1":
            forward(2) 
        if token == "[":
            stack.append([position(), heading()])
            left(45) 
        if token == "]":
            p, h = stack.pop()
            up()
            goto(p) 
            setheading(h)
            right(45)
            down()
    done()

def draw_fractal_plant(generation,angle=40,f=4):
    clearscreen() 
    speed("fastest")
    mode("logo")
    up()
    setpos(-400,-400)
    down()
    stack = []
    for token in generation:
        if token == "X":
            continue 
        if token == "F":
            forward(f) 
        if token == "+":
            left(angle) 
        if token == "-":
            right(angle)
        if token == "[":
            stack.append([position(), heading()])
        if token == "]":
            p, h = stack.pop()
            up()
            goto(p)
            setheading(h) 
            down()
    done()

def draw_fractal_plant_custom(rules,angle=40,f=5):
    gen = gen_lsystem("F", rules=rules,iterations=10, verbose=False) 
    draw_fractal_plant(gen[9],angle,f)



