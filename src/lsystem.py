def gen_lsystem(axiom, rules, iterations, verbose=True):
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