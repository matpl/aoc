import string

file = open('input.txt','r')
input = file.read().split('\n\n')
passports = list(map(lambda line : line.strip().replace('\n', ' ').split(' '), input))

def isValidHgt(h):
    qty = h[0:len(h) - 2]
    return len(h) > 2 and qty.isnumeric() and ((h[len(h) - 2:] == 'cm' and int(qty) >= 150 and int(qty) <= 193) or (h[len(h) - 2:] == 'in' and int(qty) >= 59 and int(qty) <= 76))

def isValidHcl(h):
    return len(h) == 7 and h[0] == '#' and all(c in string.hexdigits.lower() for c in h[1:])

mandatory = {
    'byr': lambda s : s.isnumeric() and len(s) == 4 and int(s) >= 1920 and int(s) <= 2002,
    'iyr': lambda s : s.isnumeric() and len(s) == 4 and int(s) >= 2010 and int(s) <= 2020,
    'eyr': lambda s : s.isnumeric() and len(s) == 4 and int(s) >= 2020 and int(s) <= 2030,
    'hgt': isValidHgt,
    'hcl': isValidHcl,
    'ecl': lambda s : s == 'amb' or s == 'blu' or s == 'brn' or s == 'gry' or s == 'grn' or s == 'hzl' or s == 'oth',
    'pid': lambda s : len(s) == 9 and s.isnumeric() }

valid = 0
for passport in passports:
    ok = True
    for m in mandatory:
        if (not any(el.startswith(m) for el in passport)) or (not mandatory[m](next(el for el in passport if el.startswith(m)).split(':')[1])):
            ok = False
            break
    if ok:
        valid += 1
print(valid)