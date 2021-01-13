file = open('input.txt','r')
input = list(map(lambda line: line.strip().split(')'),file.readlines()))

orbits = {}

for orbit in input:
    if orbit[0] not in orbits:
        orbits[orbit[0]] = None
    if orbit[1] not in orbits:
        orbits[orbit[1]] = None

    orbits[orbit[1]] = orbit[0]

def countParents(o):
    count = 0
    while orbits[o] is not None:
        count += 1
        o = orbits[o]
    return count

total = 0
for o in orbits:
    total += countParents(o)
print(total)

sanPath = []
o = 'SAN'
while orbits[o] is not None:
    o = orbits[o]
    sanPath.append(o)

transfers = 0
o = 'YOU'
while orbits[o] not in sanPath:
    transfers += 1
    o = orbits[o]
print(transfers + sanPath.index(orbits[o]))