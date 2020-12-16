import math
file = open('input.txt','r')

def parse(line):
    if line[0] != 'mask':
        line[0] = int(line[0][4:len(line[0]) - 1])
        line[1] = int(line[1])
    else:
        ormask = 0
        andmask = 0
        for i in range(len(line[1])):
            if line[1][i] == '1':
                ormask += pow(2, len(line[1]) - 1 - i)
                andmask += pow(2,len(line[1]) - 1 - i)
            elif line[1][i] == '0':
                andmask += pow(2,len(line[1]) - 1 - i)

        line[1] = [ormask, andmask, [len(line[1]) - 1 - i for i, x in enumerate(line[1]) if x == 'X']]
        #print(line[1][2])
    return line

lines = list(map(lambda line : parse(line.strip().split(' = ')), file.readlines()))

mem = {}
for line in lines:
    if line[0] == 'mask':
        masks = line[1]
    else:
        #mem[line[0]] = []

        wawa = [(line[0] | masks[0]) & masks[1]]

        #mem[line[0]].append((line[0] | masks[0]) & masks[1])
        for i in masks[2]:
            #length = len(mem[line[0]])
            length = len(wawa)
            for j in range(0, length):
                #mem[line[0]].append(mem[line[0]][j] + pow(2,i))
                wawa.append(wawa[j] + pow(2,i))
        
        for addr in wawa:
            mem[addr] = line[1]


tot = 0
for val in mem:
    tot += mem[val]
    #for innerVal in mem[val]:
    #    tot += innerVal
print(tot)
