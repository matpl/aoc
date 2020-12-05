file = open('input.txt','r')
lines = list(map(lambda line : line.strip(), file.readlines()))

total = 0
for line in lines:
    splits = line.split(':')
    numbers = splits[0].strip().split('-')
    letter = numbers[1].split(' ')[1]
    atLeast = int(numbers[0])
    atMost = int(numbers[1].split(' ')[0])
    password = splits[1].strip()
    #count = password.count(letter)
    #if(count >= atLeast and count <= atMost):
    #    total += 1

    pos1 = password[atLeast - 1] == letter
    pos2 = password[atMost - 1] == letter
    if pos1 ^ pos2:
        total += 1

print(total)