file = open('input.txt','r')
numbers = list(map(lambda c: int(c),file.read().strip().split(',')))

input = 5

line = numbers[:]
i = 0
while line[i] != 99:
    opCode = int(str(line[i])[-2:])
    parameterModes = list(map(lambda c: int(c), str(line[i])[::-1][2:]))
    while len(parameterModes) < 3:
        parameterModes.append(0)
    for j in range(0, len(parameterModes)):
        if parameterModes[j] == 0:
            if i + j + 1 < len(line) and line[i + j + 1] < len(line):
                parameterModes[j] = line[line[i + j + 1]]
        else:
            if i + j + 1 < len(line):
                parameterModes[j] = line[i + j + 1]
    
    if opCode == 1:
        line[line[i + 3]] = parameterModes[0] + parameterModes[1]
        i += 4
    elif opCode == 2:
        line[line[i + 3]] = parameterModes[0] * parameterModes[1]
        i += 4
    elif opCode == 3:
        line[line[i + 1]] = input
        i += 2
    elif opCode == 4:
        print(parameterModes[0])
        i += 2
    elif opCode == 5:
        if parameterModes[0] != 0:
            i = parameterModes[1]
        else:
            i += 3
    elif opCode == 6:
        if parameterModes[0] == 0:
            i = parameterModes[1]
        else:
            i += 3
    elif opCode == 7:
        if parameterModes[0] < parameterModes[1]:
            line[line[i + 3]] = 1
        else:
            line[line[i + 3]] = 0
        i += 4
    elif opCode == 8:
        if parameterModes[0] == parameterModes[1]:
            line[line[i + 3]] = 1
        else:
            line[line[i + 3]] = 0
        i += 4
    elif opCode == 99:
        break