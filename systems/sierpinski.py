from lsystem import gen_lsystem 

VARIABLES = ["F","G"]
CONSTANTS = ["+","-"] 
AXIOM = "F−F−F"
RULES = {"F":"F−G+F+G−F","G":"GG"} 

def gen_sierpinski(axiom=AXIOM, rules=RULES,verbose=True, iterations=6):
    return gen_lsystem(axiom, rules, iterations, verbose)

    
