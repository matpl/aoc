file = open('input.txt','r')
lines = list(map(lambda line : line.strip(), file.readlines()))

threshold = 5

def occupy(i, j):
    if lines[i][j] == 'L':
        for k in range(-1, 2):
            for l in range(-1, 2):
                if k != 0 or l != 0:
                    for m in range(0, len(lines)):
                        row = (m + 1) * k + i
                        col = (m + 1) * l + j
                        if row >= 0 and col >= 0 and row < len(lines) and col < len(lines[0]):
                            if lines[row][col] == '#':
                                return False
                            elif lines[row][col] == 'L':
                                break
                        else:
                            break
        return True
    return False

def empty(i, j):
    global threshold
    if lines[i][j] == '#':
        count = 0
        for k in range(-1, 2):
            for l in range(-1, 2):
                if k != 0 or l != 0:
                    for m in range(0, len(lines)):
                        row = (m + 1) * k + i
                        col = (m + 1) * l + j
                        if row >= 0 and col >= 0 and row < len(lines) and col < len(lines[0]):
                            if lines[row][col] == '#':
                                count += 1
                                if count >= threshold:
                                    return True
                                break
                            elif lines[row][col] == 'L':
                                break
                        else:
                            break
    return False

def iterate():
    newLines = []
    changed = False
    global lines
    for i in range(0, len(lines)):
        newLines.append([])
        for j in range(0, len(lines[i])):
            if occupy(i,j):
                newLines[i].append('#')
                changed = True
            elif empty(i,j):
                newLines[i].append('L')
                changed = True
            else:
                newLines[i].append(lines[i][j])
    lines = newLines
    return changed

while True:
    if not iterate():
        break

count = 0
for i in range(0,len(lines)):
    for j in range(0,len(lines[i])):
        if lines[i][j] == '#':
            count += 1
print(count)