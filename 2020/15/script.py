import math
file = open('input.txt','r')

lines = list(map(lambda line : line.strip(), file.readlines()))
line = list(map(lambda n: int(n), lines[0].split(',')))

occ = {}
for i in range(0, len(line)):
    occ[line[i]] = [i]

#pos = 2020
pos = 30000000
for i in range(0, pos - len(line)):
    if len(occ[line[-1]]) == 1:
        val = 0
    else:
        val = occ[line[-1]][-1] - occ[line[-1]][-2]

    if not val in occ:
        occ[val] = []
    occ[val].append(len(line))
    line.append(val)

print(line[-1])
        