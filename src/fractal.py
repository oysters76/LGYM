from assert_utils import assert_array 
'''
 variables: 0, 1
 constants: "[", "]" 
 axiom: 0 
 rules: (1 -> 11), (0 -> 1[0]0) 
'''
VARIABLES = ["0","1"] 
CONSTANTS = ["[", "]"] 
AXIOM = "0" 
RULES = {"1":"11", "0":"1[0]0"}

def gen_fractal(variables=VARIABLES, 
                constants=CONSTANTS, axiom=AXIOM, 
                rules=RULES,verbose=True, iterations=4):
    next_generation = "" 
    current_generation = axiom 
    generations = [] 
    for n in range(iterations):
        for token in current_generation:

            if token not in rules:
                next_generation += token 
                continue 

            next_generation += rules[token]

        
        if verbose:
            print("n = ", n, current_generation) 
        
        generations.append(current_generation)
        current_generation = next_generation 
        next_generation = ""

    return generations


def assert_gen_fractal():
    expected = [
        "0",
        "1[0]0",
        "11[1[0]0]1[0]0",
        "1111[11[1[0]0]1[0]0]11[1[0]0]1[0]0"
    ]
    actual = gen_fractal(verbose=False)
    assert_array(actual, expected)


assert_gen_fractal()