file = open('input.txt','r')
lines = list(map(lambda line : list(line.strip()), file.readlines()))

ROW_RANGE = list(range(0,128))
COL_RANGE = list(range(0,8))

seats = []

for line in lines:
    row = ROW_RANGE[:]
    col = COL_RANGE[:]
    for i in range(0,7):
        if line[i] == 'F':
            row = row[:int(len(row)/2)]
        else:
            row = row[int(len(row)/2):]
    for i in range(7,10):
        if line[i] == 'L':
            col = col[:int(len(col)/2)]
        else:
            col = col[int(len(col)/2):]
    seats.append(row[0] * 8 + col[0])

seats.sort()
for i in range(0, len(seats) - 1):
    if seats[i] + 1 != seats[i + 1]:
        print(seats[i] + 1)
        break
