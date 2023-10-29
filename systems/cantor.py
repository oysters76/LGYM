from lsystem import gen_lsystem 

VARIABLES = ["A","B"]
CONSTANTS = [] 
AXIOM = "A" 
RULES = {"A":"ABA", "B":"BBB"} 

def gen_cantor(axiom,rules=RULES, iterations=4):
    return gen_lsystem(axiom, rules, iterations) 

