file = open('input.txt','r')
lines = list(map(lambda line : int(line.strip()), file.readlines()))

def isValid(arr, numba):
    for i in range(0, len(arr)):
        for j in range(0, len(arr)):
            if i != j and arr[i] + arr[j] == numba:
                return True
    return False

for i in range(25,len(lines)):
    if not isValid(lines[i-25:i], lines[i]):
        invalid = lines[i]
        break

for i in range(0,len(lines)):
    sum = lines[i]
    for j in range(i+1,len(lines)):
        sum += lines[j]
        if sum == invalid:
            start = i
            end = j
            break
        elif sum > invalid:
            break

minV = invalid
maxV = 0
for i in range(start,end):
    minV = min(minV, lines[i])
    maxV = max(maxV,lines[i])

print(minV + maxV)