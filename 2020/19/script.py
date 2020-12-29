import math
file = open('input.txt','r')

lines = list(map(lambda line : line.strip().replace('"',''), file.readlines()))
messages = []
rules = {}
for line in lines:
    if line != '':
        if line[0].isnumeric():
            splits = line.split(': ')
            groups = splits[1].split(' | ')
            for i in range(0, len(groups)):
                groups[i] = groups[i].split(' ')
            rules[int(splits[0])] = groups
        else:
            messages.append(line)

rules[8] = [['42'], ['42', '8']]
rules[11] = [['42', '31'], ['42', '11', '31']]

for rule in rules:
    for i in range(0, len(rules[rule])):
        for j in range(0, len(rules[rule][i])):
            if rules[rule][i][j].isnumeric():
                rules[rule][i][j] = rules[int(rules[rule][i][j])]

def visit(rule, msg, pos):
    valid = []
    for p in pos:
        if p == len(msg):
            continue
        if rule[0][0] == 'a':
            if msg[p] == 'a':
                if p + 1 not in valid:
                    valid.append(p + 1)
                continue
            else:
                continue
        if rule[0][0] == 'b':
            if msg[p] == 'b':
                if p + 1 not in valid:
                    valid.append(p + 1)
                continue
            else:
                continue
        for r in rule: # or
            newPos = pos
            for innerR in r: # and
                newPos = visit(innerR, msg, newPos)
                if len(newPos) == 0:
                    break
            for np in newPos:
                if np not in valid:
                    valid.append(np)
    return valid

count = 0
for msg in messages:
    res = visit(rules[0], msg, [0])
    for r in res:
        if r == len(msg):
            count += 1
            break
print(count)