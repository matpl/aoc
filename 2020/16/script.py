import math
file = open('input.txt','r')

lines = list(map(lambda line : line.strip(), file.readlines()))

ranges = []
nearby = []
i = 0
while "your ticket" not in lines[i]:
    if lines[i] != '':
        r = lines[i].split(': ')
        valid = list(map(lambda num : num.split('-'), r[1].split(' or ')))
        ranges.append([r[0], list(map(lambda num : [int(num[0]), int(num[1])], valid))])
    i += 1

ticket = list(map(lambda line: int(line), lines[i + 1].split(',')))

for i in range(i + 4, len(lines)):
    nearby.append(list(map(lambda line : int(line), lines[i].split(','))))

def isValid(n):
    for r in ranges:
        for r2 in r[1]:
            if n >= r2[0] and n <= r2[1]:
                return True
    return False

wrong = []
for i in range(0, len(nearby)):
    for n in nearby[i]:
        if not isValid(n):
            if i not in wrong:
                wrong.append(i)
wrong.reverse()
for i in wrong:
    del nearby[i]

def allValid(pos, j):
    for i in range(0, len(pos)):
        fits = False
        for r in ranges[j][1]:
            if pos[i] >= r[0] and pos[i] <= r[1]:
                fits = True
                break
        if not fits:
            return False
    return True

possibilities = {}
for i in range(0, len(ranges)):
    pos = []
    for t in nearby:
        pos.append(t[i])
    for j in range(0, len(ranges)):
        if allValid(pos, j):
            if ranges[j][0] not in possibilities:
                possibilities[ranges[j][0]] = []
            possibilities[ranges[j][0]].append(i)

possibilities = list(map(lambda pos: [pos, possibilities[pos]], possibilities))
possibilities.sort(key=lambda x: len(x[1]))
positions = {}
while len(possibilities) != 0:
    positions[possibilities[0][0]] = possibilities[0][1][0]
    for i in range(1, len(possibilities)):
        possibilities[i][1].remove(possibilities[0][1][0])
    del possibilities[0]
mult = 1
for pos in positions:
    if 'departure' in pos:
        mult *= ticket[positions[pos]]
print(mult)