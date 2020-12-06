file = open('input.txt','r')
input = list(map(lambda c: int(c),file.read().strip().split(',')))
output = 19690720
for j in range(0, 100):
    for k in range(0, 100):
        line = input[:]
        line[1] = j
        line[2] = k
        i = 0
        while line[i] != 99:
            if line[i] == 1:
                line[line[i + 3]] = line[line[i + 1]] + line[line[i + 2]]
                i += 4
            elif line[i] == 2:
                line[line[i + 3]] = line[line[i + 1]] * line[line[i + 2]]
                i += 4
            elif line[i] == 99:
                break
        if(line[0] == output):
            print(100 * j + k)