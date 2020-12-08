file = open('input.txt','r')

def parseLine(line):
    line[1] = int(line[1])
    return line

lines = list(map(lambda line : parseLine(line.strip().split(' ')), file.readlines()))

def visit():
    visited = {}
    accumulator = 0
    i = 0
    while i < len(lines):
        if lines[i][0] == 'acc':
            if not i in visited:
                visited[i] = True
            else:
                return
            accumulator += lines[i][1]
        elif lines[i][0] == 'jmp':
            i += lines[i][1] - 1
        i += 1
    print('finished: ' + str(accumulator))

for i in range(len(lines)):
    if lines[i][0] == 'nop':
        lines[i][0] = 'jmp'
        visit()
        lines[i][0] = 'nop'
    elif lines[i][0] == 'jmp':
        lines[i][0] = 'nop'
        visit()
        lines[i][0] = 'jmp'