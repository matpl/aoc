import math
file = open('input.txt','r')

def parse(line):
    return line.strip()

lines = list(map(lambda line : parse(line), file.readlines()))

buses = lines[1].split(',')
for i in range(0, len(buses)):
    if buses[i] != 'x':
        buses[i] = [int(buses[i]), i]
buses = [bus for bus in buses if bus != 'x']
offset = buses[0][1]
mult = 1
for i in range(0, len(buses)):
    mult = mult * buses[i][0]
    buses[i][1] = buses[i][1] - offset

no = 0
inc = buses[0][0]
i = 1
while i < len(buses) :
    if (no + buses[i][1]) % buses[i][0] == 0:
        inc = inc * buses[i][0]
        i += 1
        if i == len(buses):
            print(no)
    no += inc