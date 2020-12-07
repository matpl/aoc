file = open('input.txt','r')

def parseLine(line):
    line = line.strip().replace('.','').replace(' bags','').replace(' bag','').replace('no','0').split(' contain ')
    line[1] = line[1].split(', ')
    for i in range(len(line[1])):
        line[1][i] = [int(line[1][i][:line[1][i].index(' ')]),line[1][i][line[1][i].index(' ') + 1:]]
    return line

lines = list(map(lambda line : parseLine(line), file.readlines()))

bags = {}
for line in lines:
    bags[(1,line[0])] = []
for line in lines:
    for content in line[1]:
        if content[0] != 0:
            bags[(1,line[0])].append((content[0],bags[(1,content[1])]))

shiny = bags[(1,'shiny gold')]
def rec(arr):
    if arr is shiny:
        return 1
    else:
        for i in arr:
            if rec(i[1]):
                return 1
        return 0

count = 0
for k in bags:
    count += rec(bags[k])
print(count - 1)

def bagCount(arr):
    count = 0
    for i in arr:
        count += i[0] + i[0] * bagCount(i[1])
    return count

print(bagCount(shiny))