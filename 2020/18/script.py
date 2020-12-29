import math
file = open('input.txt','r')

lines = list(map(lambda line : list(line.strip().replace(' ','')), file.readlines()))

def parse(line):
    arr = []
    while len(line) != 0:
        if line[0] == '(':
            line.pop(0)
            arr.append(parse(line))
        elif line[0] == ')':
            line.pop(0)
            break
        else:
            arr.append(line.pop(0))
    return arr


def calc(form):
    for i in range(0, len(form)):
        if isinstance(form[i], list):
            form[i] = calc(form[i])
    #while len(form) != 1:
    #    if form[1] == '+':
    #        form[0] = int(form.pop(2)) + int(form.pop(0))
    #    elif form[1] == '*':
    #        form[0] = int(form.pop(2)) * int(form.pop(0))
    for i in reversed(range(0, len(form) - 2)):
        if form[i + 1] == '+':
            form[i] = int(form.pop(i + 2)) + int(form.pop(i))
    for i in reversed(range(0, len(form) - 2)):
        if form[i + 1] == '*':
            form[i] = int(form.pop(i + 2)) * int(form.pop(i))
    return form[0]
            
tot = 0
for line in lines:
    tot += int(calc(parse(line)))
print(tot)