from lsystem import gen_lsystem 

AXIOM = "F" 
RULES = {"F":"F+G", "G":"F-G"}

def gen_dragon_curve(axiom=AXIOM, rules=RULES,verbose=True, iterations=10):
    return gen_lsystem(axiom, rules, iterations, verbose)
