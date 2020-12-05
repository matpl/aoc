import math

file = open('input.txt','r')

def fuel(n, i = 0):
    f = max(math.floor(n / 3) - 2, 0)
    if i == 0:
        n = 0
    return n + (f if f == 0 else fuel(f, i + 1))

lines = list(map(lambda line : fuel(int(line.strip())), file.readlines()))
print(sum(lines))