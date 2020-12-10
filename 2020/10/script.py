file = open('input.txt','r')
lines = list(map(lambda line : int(line.strip()), file.readlines()))
lines.sort()
lines.insert(0,0)
totalsAt = {}
def check(index):
    if index in totalsAt:
        return totalsAt[index]
    total = 0
    for i in range(index + 1, len(lines)):
        if lines[i] - lines[index] > 3:
            break
        elif lines[i] == lines[-1]:
            total += 1
        else:
            total += check(i)
    totalsAt[index] = total
    return total
print(check(0))