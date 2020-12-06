file = open('input.txt','r')

def inc(letter):
    if letter == 'U':
        return (1,0)
    elif letter == 'D':
        return (-1,0)
    elif letter == 'R':
        return (0,1)
    elif letter == 'L':
        return (0,-1)

def extract(r):
    splits = r.strip().split(',')
    return list(map(lambda split: [inc(split[0]), int(split[1:])], splits))

input = list(map(lambda c: extract(c), file.readlines()))

covered = {}
for j in range(0,len(input)):
    line = input[j]
    pos = (0, 0)
    steps = 0
    for move in line:
        for i in range(0, move[1]):
            steps += 1
            pos = (pos[0] + move[0][0], pos[1] + move[0][1])
            if pos in covered:
                if covered[pos][j] == 0:
                    if j == 0:
                        covered[pos] = (steps, covered[pos][1])
                    else:
                        covered[pos] = (covered[pos][0], steps)
            else:
                if j == 0:
                    covered[pos] = (steps, 0)
                else:
                    covered[pos] = (0, steps)

closest = float('inf')
minSteps = float('inf')

for k in covered:
    if covered[k][0] != 0 and covered[k][1] != 0:
        closest = min(closest, abs(k[0]) + abs(k[1]))
        minSteps = min(minSteps, covered[k][0] + covered[k][1])

print(closest)
print(minSteps)