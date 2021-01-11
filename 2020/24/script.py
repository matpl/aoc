file = open('input.txt','r')

def parseLine(line):
    i = 0
    tile = (0, 0)
    while i != len(line):
        if line[i] == 'n': # north
            if line[i + 1] == 'w': # west
                tile = (tile[0] - 0.5, tile[1] - 0.5)
            else: # east
                tile = (tile[0] - 0.5, tile[1] + 0.5)
            i += 2
        elif line[i] == 's': # south
            if line[i + 1] == 'w': # west
                tile = (tile[0] + 0.5, tile[1] - 0.5)
            else: # east
                tile = (tile[0] + 0.5, tile[1] + 0.5)
            i += 2
        else:
            if line[i] == 'w':
                tile = (tile[0], tile[1] - 1)
            elif line[i] == 'e':
                tile = (tile[0], tile[1] + 1)
            i += 1
    return tile

lines = list(map(lambda line : parseLine(line.strip()), file.readlines()))
tiles = {}
for line in lines:
    if line not in tiles:
        tiles[line] = True
    else:
        tiles[line] = not tiles[line]

def countBlacks():
    blacks = 0
    for tile in tiles:
        if tiles[tile]:
            blacks += 1
    return blacks
print(countBlacks())

def getNeighbors(tile):
    return [
        (tile[0] - 0.5, tile[1] - 0.5),
        (tile[0] - 0.5, tile[1] + 0.5),
        (tile[0], tile[1] + 1),
        (tile[0] + 0.5, tile[1] + 0.5),
        (tile[0] + 0.5, tile[1] - 0.5),
        (tile[0], tile[1] - 1)
    ]

def countBlackNeighbors(tile):
    neighbors = getNeighbors(tile)
    blackNeighbors = 0
    for neighbor in neighbors:
        if neighbor in tiles and tiles[neighbor]:
            blackNeighbors += 1
    return blackNeighbors

def appendWhiteNeighbors(tile):
    neighbors = getNeighbors(tile)
    for neighbor in neighbors:
        if neighbor not in tiles:
            tiles[neighbor] = False

for i in range(0, 100):
    for tile in list(tiles.keys()):
        if tiles[tile]:
            appendWhiteNeighbors(tile)
    toFlip = []
    for tile in tiles:
        blackNeighbors = countBlackNeighbors(tile)
        if tiles[tile]:
            if blackNeighbors == 0 or blackNeighbors > 2:
                toFlip.append(tile)
        else:
            if blackNeighbors == 2:
                toFlip.append(tile)
    for tile in toFlip:
        tiles[tile] = not tiles[tile]
print(countBlacks())
