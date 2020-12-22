import math
file = open('input.txt','r')

lines = list(map(lambda line : line.strip(), file.readlines()))

cubesize = 25

grid = []
for i in range(0,cubesize):
    grid.append([])
    for j in range(0,cubesize):
        grid[i].append([])
        for k in range(0,cubesize):
            grid[i][j].append([])
            for l in range(0,cubesize):
                grid[i][j][k].append('.')

x = int(cubesize / 2)
z = int(cubesize / 2)
w = int(cubesize / 2)
for line in lines:
    y = int(cubesize / 2)
    for c in line:
        grid[x][y][z][w] = c
        y += 1
    x += 1

def activeNeighbors(i, j, k, l):
    count = 0
    for x in range(-1,2):
        for y in range(-1,2):
            for z in range(-1,2):
                for w in range(-1,2):
                    if x != 0 or y != 0 or z != 0 or w != 0:
                        if grid[i+x][j+y][k+z][l+w] == '#':
                            count += 1
    return count

def flip(i, j, k, l):
    count = activeNeighbors(i, j, k, l)

    if grid[i][j][k][l] == '#':
        if count == 2 or count == 3:
            return False
        else:
            return True
    elif grid[i][j][k][l] == '.':
        if count == 3:
            return True
    return False

for m in range(0, 6):
    toflip = []
    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[i]) - 1):
            for k in range(1, len(grid[i][j]) - 1):
                for l in range(1, len(grid[i][j][k]) - 1):
                    if flip(i, j, k, l):
                        toflip.append((i, j, k, l))
    for f in toflip:
        if grid[f[0]][f[1]][f[2]][f[3]] == '#':
            grid[f[0]][f[1]][f[2]][f[3]] = '.'
        else:
            grid[f[0]][f[1]][f[2]][f[3]] = '#'

count = 0
for i in range(0, len(grid)):
    for j in range(0, len(grid[i])):
        for k in range(0, len(grid[i][j])):
            for l in range(0, len(grid[i][j][k])):
                if grid[i][j][k][l] == '#':
                    count += 1
print(count)