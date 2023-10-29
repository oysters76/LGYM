from lsystem import gen_lsystem 
from assert_utils import assert_array

VARIABLES = ["F"] 
CONSTANTS = ["+", "−"] 
AXIOM = "F" 
RULES = {"F":"F+F−F−F+F"} 

def gen_koch(axiom=AXIOM, rules=RULES, iterations=4, verbose=True):
    return gen_lsystem(axiom, rules, iterations, verbose) 

def assert_koch():
    expected = [
        "F",
        "F+F−F−F+F", 
        "F+F−F−F+F+F+F−F−F+F−F+F−F−F+F−F+F−F−F+F+F+F−F−F+F",
        "F+F−F−F+F+F+F−F−F+F−F+F−F−F+F−F+F−F−F+F+F+F−F−F+F+F+F−F−F+F+F+F−F−F+F−F+F−F−F+F−F+F−F−F+F+F+F−F−F+F−F+F−F−F+F+F+F−F−F+F−F+F−F−F+F−F+F−F−F+F+F+F−F−F+F−F+F−F−F+F+F+F−F−F+F−F+F−F−F+F−F+F−F−F+F+F+F−F−F+F+F+F−F−F+F+F+F−F−F+F−F+F−F−F+F−F+F−F−F+F+F+F−F−F+F"
    ]
    actual = gen_koch(verbose=False) 
    assert_array(actual, expected)

assert_koch()