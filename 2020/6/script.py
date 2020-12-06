file = open('input.txt','r')
groups = list(map(lambda line : line, file.read().split('\n\n')))
count = 0
for group in groups:
    people = group.split('\n')
    questions = (1 << 26) - 1
    for p in people:
        p = list(map(lambda c: 1 << (ord(c) - 97), p))
        bin = 0
        for c in p:
            bin |= c
        questions &= bin

    for i in range(0, 26):
        count += (questions >> i) & 1

print(count)