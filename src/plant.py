from lsystem import gen_lsystem 

AXIOM = "X" 
RULES = {"X":"F+[[X]-X]-F[-FX]+X", "F":"FF"}

def gen_fractal_plant(axiom=AXIOM, rules=RULES, verbose=True, iterations=16):
    return gen_lsystem(axiom, rules, iterations, verbose) 


