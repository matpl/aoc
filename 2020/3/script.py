file = open('input.txt','r')
lines = list(map(lambda line : list(line.strip()), file.readlines()))

def treeCount(offsetX, offsetY):
    x = 0
    y = 0
    count = 0
    while y < len(lines):
        if lines[y][x] == '#':
            count += 1
        x = (x + offsetX) % len(lines[0])
        y += offsetY

    return count

print(treeCount(1,1) * treeCount(3,1) * treeCount(5,1) * treeCount(7,1) * treeCount(1,2))