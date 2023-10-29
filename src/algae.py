from assert_utils import assert_array
'''
 variables A, B 
 constants None 
 axiom     A 
 rules (A -> AB), (B -> A)
'''

def gen_algae(axiom,rules={"A":"AB","B":"A"}, verborse=True, iterations=8):
    next_generation = "" 
    current_generation = axiom 
    generations = [] 

    for n in range(0, iterations):
        for i in current_generation:
            if i in rules:
                next_generation += rules[i] 
            else:
                next_generation += i 
        
        if verborse:
            print("n = ", n, current_generation)

        generations.append(current_generation)
        current_generation = next_generation
        next_generation = ""
    
    return generations

def assert_gen_algae():
    expected = [
        "A",
        "AB",
        "ABA",
        "ABAAB",
        "ABAABABA",
        "ABAABABAABAAB",
        "ABAABABAABAABABAABABA",
        "ABAABABAABAABABAABABAABAABABAABAAB",
    ] 
    actual = gen_algae("A", verborse=False)
    assert_array(actual, expected)

assert_gen_algae()

