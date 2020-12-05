file = open('input.txt','r')
numbers = list(map(lambda line : int(line.strip()), file.readlines()))

RESULT = 2020

for i in range(0,len(numbers)):
    for j in range(i + 1, len(numbers)):
        for k in range(j + 1, len(numbers)):
            if numbers[i] + numbers[j] + numbers[k] == RESULT:
                print(numbers[i] * numbers[j] * numbers[k])