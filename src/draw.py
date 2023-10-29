
from turtle import * 

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

