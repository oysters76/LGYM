from lsystem import gen_lsystem 

VARIABLES = ["F","G"]
CONSTANTS = ["+","-"] 
AXIOM = "F−F−F"
RULES = {"F":"F−G+F+G−F","G":"GG"} 

def gen_sierpinski(axiom=AXIOM, rules=RULES,verbose=True, iterations=6):
    return gen_lsystem(axiom, rules, iterations, verbose)


A_AXIOM = "A"
A_RULES = {"A":"B−A−B", "B":"A+B+A"}

def gen_sierpinski_arrowhead(axiom=A_AXIOM, rules=A_RULES, verbose=True, iterations=6):
    return gen_lsystem(axiom, rules, iterations, verbose)

