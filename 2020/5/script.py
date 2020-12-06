file = open('input.txt','r')
lines = list(map(lambda line : list(line.strip()), file.readlines()))
seats = []

for line in lines:
    pos = 1023
    for i in range(0,10):
        if line[i] == 'F' or line[i] == 'L':
            pos &= ~(512 >> i)
    seats.append(pos)

seats.sort()
for i in range(0, len(seats) - 1):
    if seats[i] + 1 != seats[i + 1]:
        print(seats[i] + 1)
        break