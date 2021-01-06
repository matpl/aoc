file = open('input.txt','r')
lines = list(map(lambda line : line.strip()[:-1].split(' ('), file.readlines()))
allergens = {}
for line in lines:
    line[0] = line[0].split(' ')
    line[1] = line[1].replace(',','').replace('contains ', '').split(' ')
    for allergen in line[1]:
        if allergen not in allergens:
            allergens[allergen] = []
        allergens[allergen].append([])
        for ing in line[0]:
            allergens[allergen][-1].append(ing)

solved = {}
found = True
while found:
    found = False
    for allergen in list(allergens.keys()):
        if len(allergens[allergen]) > 1:
            # if only one ingredient is there in all lists, it's the allergen
            many = []
            for ing in allergens[allergen][0]:
                total = 1
                for i in range(1, len(allergens[allergen])):
                    total += allergens[allergen][i].count(ing)
                if total == len(allergens[allergen]):
                    many.append(ing)
            if len(many) == 1:
                found = True
                solved[allergen] = many[0]
                # remove this ingredient from other allergens
                for otherAllergen in allergens:
                    for r in allergens[otherAllergen]:
                        if solved[allergen] in r:
                            r.remove(solved[allergen])
                del allergens[allergen]
        elif len(allergens[allergen][0]) == 1:
            found = True
            solved[allergen] = allergens[allergen][0][0]
            del allergens[allergen]

total = 0
for line in lines:
    count = 0
    for v in solved.values():
        count += line[0].count(v)
    total += len(line[0]) - count
print(total)

keys = list(solved.keys())
keys.sort()
dangerous = ''
for k in keys:
    if dangerous != '':
        dangerous += ','
    dangerous += solved[k]
print(dangerous)