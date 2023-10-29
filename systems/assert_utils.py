def assert_array(actual, expected):
    for i, row in enumerate(actual):
        for j, l in enumerate(row):
            expected_row = expected[i] 
            assert l == expected_row[j], "row " + str(i) + " failed: expected: " + expected_row + "\t actual: " + row 