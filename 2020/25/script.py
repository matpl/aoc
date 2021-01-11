file = open('input.txt','r')
lines = list(map(lambda line : int(line.strip()), file.readlines()))

def transform(subject, loopSize):
    return pow(subject, loopSize, 20201227)

def findLoopSize(publicKey):
    loopSize = 1
    while True:
        v = transform(7, loopSize)
        if v == publicKey:
            return loopSize
        loopSize += 1

print(transform(lines[1], findLoopSize(lines[0])))

# Or simply:
# https://www.wolframalpha.com/input/?i=solve+%287+%5E+n+-+9717666%29+mod+20201227+%3D+0+for+n
# https://www.wolframalpha.com/input/?i=20089533+%5E+17167199+mod+20201227
# ;)
