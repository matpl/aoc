import math
file = open('input.txt','r')

def parse(line):
    c = line.strip()[0]
    if c == 'N':
        c = (0,1)
    elif c == 'S':
        c = (0,-1)
    elif c == 'E':
        c = (1,0)
    elif c == 'W':
        c = (-1,0)
    
    return [c, int(line.strip()[1:])]

lines = list(map(lambda line : parse(line), file.readlines()))

def rotate(deg):
    global x, y
    oldX = x
    x = round(x * math.cos(deg * math.pi / 180) - y * math.sin(deg * math.pi / 180))
    y = round(oldX * math.sin(deg * math.pi / 180) + y * math.cos(deg * math.pi / 180))

#dirs = [(1,0), (0,-1), (-1,0), (0,1)]
x = 10
y = 1
shipX = 0
shipY = 0
dir = 0
for line in lines:
    if line[0] == 'F':
        #x += dirs[dir][0] * line[1]
        #y += dirs[dir][1] * line[1]
        shipX += line[1] * x
        shipY += line[1] * y
    elif line[0] == 'L':
        #dir = (dir - int(line[1] / 90)) % len(dirs)
        rotate(line[1])
    elif line[0] == 'R':
        #dir = (dir + int(line[1] / 90)) % len(dirs)
        rotate(line[1] * -1)
    else:
        x += line[0][0] * line[1]
        y += line[0][1] * line[1]

print(abs(shipX) + abs(shipY))