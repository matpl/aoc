file = open('input.txt','r')
lines = list(map(lambda line : list(line.strip()), file.readlines()))
seats = []

for line in lines:
    row = 127
    col = 7
    for i in range(0,7):
        if line[i] == 'F':
            row &= ~(64 >> i)
    for i in range(7,10):
        if line[i] == 'L':
            col &= ~(4 >> (i - 7))
    seats.append(row * 8 + col)

seats.sort()
for i in range(0, len(seats) - 1):
    if seats[i] + 1 != seats[i + 1]:
        print(seats[i] + 1)
        break
