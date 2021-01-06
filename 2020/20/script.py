import math
file = open('input.txt','r')

tilesize = 10
tiles = {}
tileBorders = {}

for line in file.readlines():
    line = line.strip()
    if line.startswith('Tile'):
        tile = []
        tiles[int(line.replace(':','').split(' ')[1])] = tile
    elif line != '':
        tile.append(line)

for tile in tiles:
    values = [['','','',''], ['','','','']]
    for i in range(0, tilesize):
        values[0][0] += tiles[tile][0][i] # top
        values[0][1] += tiles[tile][i][tilesize - 1] # right
        values[0][2] += tiles[tile][tilesize - 1][tilesize - 1 - i] # bottom
        values[0][3] += tiles[tile][tilesize - 1 - i][0] # left
    for j in reversed(range(0, tilesize)):
        values[1][2] += values[0][2][j]
        values[1][3] += values[0][1][j]
        values[1][0] += values[0][0][j]
        values[1][1] += values[0][3][j]

    tileBorders[tile] = values

def uniqueborder(b):
    count = 0
    for tile in tileBorders:
        for piece in tileBorders[tile]:
            for border in piece:
                if b == border:
                    count += 1
                    if count > 1:
                        return False
    return True

def rotate(tile, rotations, flipped):
    if flipped == 1: # flip the piece before rotation
        flippedTile = []
        for i in range(0, len(tiles[tile])):
            flippedTile.append('')
            for j in reversed(range(0, len(tiles[tile][i]))):
                flippedTile[-1] += tiles[tile][i][j]
        tiles[tile] = flippedTile

    for i in range(0, rotations):
        rotated = []
        for j in range(0, len(tiles[tile])):
            row = ''
            for k in reversed(range(0, len(tiles[tile]))):
                row += tiles[tile][k][j]
            rotated.append(row)
        tiles[tile] = rotated
    rotated = ['', '', '', '']
    for j in range(0, len(tileBorders[tile][flipped])):
        rotated[(j + rotations) % 4] = tileBorders[tile][flipped][j]
    tileBorders[tile] = [rotated]

solved = [[]]
# find first piece with 4 uniques borders: it's a corner
for tile in tileBorders:
    uniques = []
    for i in range(0, len(tileBorders[tile][0])):
        if uniqueborder(tileBorders[tile][0][i]):
            uniques.append(i)
    
    if len(uniques) == 2:
        # this is the top left corner
        if uniques[0] == 0 and uniques[1] == 3:
            uniques[0] = 3
            uniques[1] = 0
        
        rotations = 3 - uniques[0] # needed clockwise rotations to fit the top left corner
        rotate(tile, rotations, 0)
        solved[0].append(tile)
        break

width = int(pow(len(tiles), 0.5))

for i in range(0, width):
    for j in range(0, width):
        if i != 0 or j != 0:
            # find the right piece that goes here
            if j != 0: # compare with the piece to the left to find which one matches its right border
                referenceTile = solved[i][j - 1]
                border = tileBorders[referenceTile][0][1]
                referenceBorder = 3
            else:
                solved.append([])
                referenceTile = solved[i - 1][j]
                border = tileBorders[referenceTile][0][2]
                referenceBorder = 4

            for tile in tileBorders:
                if tile != referenceTile:
                    for k in range(0, len(tileBorders[tile])):
                        for l in range(0, len(tileBorders[tile][k])):
                            if tileBorders[tile][k][l][::-1] == border:

                                rotate(tile, referenceBorder - l, k)

                                solved[i].append(tile)
                                break
                        if tile in solved[i]:
                            break

print(solved[0][0] * solved[0][width - 1] * solved[width - 1][0] * solved[width - 1][width - 1])

image = []
for row in solved:
    for i in range(1, tilesize - 1):
        image.append('')
        for tile in row:
            image[-1] += tiles[tile][i][1:len(tiles[tile][i]) - 1]

patterns = [['                  # ', '#    ##    ##    ###', ' #  #  #  #  #  #   ']]
for i in range(0, 3): # rotate the pattern
    reference = patterns[-1] # last column becomes firt row
    patterns.append([])
    for j in reversed(range(0, len(reference[0]))):
        row = ''
        for k in range(0, len(reference)):
            row += reference[k][j]
        patterns[-1].append(row)
        
for i in range(0, len(patterns)): # flip them
    patterns.append([])
    for row in patterns[i]:
        patterns[-1].append(row[::-1])

monsterCount = 0
for p in patterns: # try to find the pattern
    for i in range(0, len(image) - len(p)):
        for j in range(0, len(image[0]) - len(p[0])):
            found = True
            for r in range(0, len(p)):
                for c in range(0, len(p[0])):
                    if p[r][c] == '#' and image[i + r][j + c] != '#':
                        found = False
            if found:
                monsterCount += 1

charCount = 0
for r in image:
    for c in r:
        if c == '#':
            charCount += 1
print(charCount - 15 * monsterCount) # 15 hashtags in a monster