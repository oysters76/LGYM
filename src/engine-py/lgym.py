'''
    Engine that parses and converts l-system to a render program
'''
import json
from turtle import * 
from lsystem import gen_lsystem 

FORWARD = "forward" 
LEFT    = "left"
RIGHT   = "right"
POP     = "pop"
PUSH    = "push"


class LConfig:
    def __init__(self):
        self.title = "" 
        self.rules = {} 
        self.cmds = {} 
        self.axiom = "" 
        self.magnitude = 0 
        self.angle = 0 
        self.ix = 0 
        self.iy = 0 
        self.width = 0 
        self.height = 0 
        self.iangle = 0
        self.iterations = 0

def _get_safe(data, key):
    if key in data:
        return data[key] 
    return ""

def _get_safe_int(data, key):
    try:
        iv = int(_get_safe(data, key)) 
        return iv
    except:
        raise RuntimeError("[ERROR]: invalid value for field ", key)

def _get_safe_dict(data, key):
    d =_get_safe(data, key) 
    if d == "" or type(d) != dict:
        raise RuntimeError("[ERROR]: invalid value for field <map> ", key) 
    return d 

def parse(filename : str) -> LConfig:
    with open(filename, "r") as input_file:
        lconfig = LConfig() 
        data = json.load(input_file) 
        lconfig.title = _get_safe(data, "title")
        lconfig.axiom = _get_safe(data, "axiom") 
        lconfig.iterations = _get_safe_int(data, "iterations")
        lconfig.magnitude = _get_safe_int(data, "magnitude") 
        lconfig.angle = _get_safe_int(data, "angle") 
        lconfig.ix = _get_safe_int(data, "ix") 
        lconfig.iy = _get_safe_int(data, "iy") 
        lconfig.width = _get_safe_int(data, "width")
        lconfig.height = _get_safe_int(data, "height") 
        lconfig.rules = _get_safe_dict(data, "rules") 
        lconfig.cmds = _get_safe_dict(data, "cmds")
        return lconfig
        
def gen_lsystem(axiom, rules, iterations, verbose=True) -> str:
    next_generation = "" 
    current_generation = axiom 
    for n in range(iterations):
        for token in current_generation:

            if token not in rules:
                next_generation += token 
                continue 

            next_generation += rules[token]

        
        if verbose:
            print("n = ", n, current_generation) 
        
        current_generation = next_generation 
        next_generation = ""
    return current_generation

def to_lsystem(lconfig : LConfig) -> str:
    return gen_lsystem(lconfig.axiom, lconfig.rules, lconfig.iterations, False) 

def to_program(lconfig : LConfig, lsys :str) -> list:
    prog = [] 
    for v in lsys:
        if v in lconfig.cmds:
            prog.append(lconfig.cmds[v]) 
    return prog

def render(prog : list, lconfig : LConfig, speedt="fastest", modet="logo") -> None:
    title(lconfig.title)
    screensize(lconfig.width, lconfig.height)
    clearscreen() 
    speed(speedt)
    mode(modet)
    up()
    setpos(lconfig.ix,lconfig.iy)
    down()
    stack = [] 
    for command in prog:
        if command == FORWARD:
            forward(lconfig.magnitude) 
        if command == LEFT:
            left(lconfig.angle) 
        if command == RIGHT:
            right(lconfig.angle)
        if command == PUSH:
            stack.append([position(), heading()])
        if command == POP:
            p, h = stack.pop()
            up()
            goto(p)
            setheading(h) 
            down()
    done()

def run_lsys(filename):
    lconfig = parse(filename) 
    lsys = to_lsystem(lconfig)
    prog = to_program(lconfig, lsys)
    render(prog, lconfig)

run_lsys("./examples/ftree.json")