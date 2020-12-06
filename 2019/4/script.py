passwords = open('input.txt','r').read().strip().split('-')
passwords = list(map(lambda x : int(x), passwords))

count = 0
for i in range(passwords[0], passwords[1] + 1):
    digits = list(map(lambda d: int(d) , str(i)))
    valid = False
    for j in range(0, len(digits) - 1):
        if digits[j] == digits[j + 1]:
            if not (j > 0 and digits[j - 1] == digits[j]) and not (j < len(digits) - 2 and digits[j + 2] == digits[j]):
                valid = True
        if digits[j] > digits[j + 1]:
            valid = False
            break
    if valid:
        count += 1
print(count)