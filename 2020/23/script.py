from collections import deque

input = "467528193"

class Node:
    def __init__(self, number):
        self.number = number
        self.next = None

indexes = []
for i in range(0, 1000000):
    node = Node(i + 1)
    if i != 0:
        indexes[-1].next = node
    indexes.append(node)
for i in range(0, len(input)):
    if i != len(input) - 1:
        indexes[int(input[i]) - 1].next = indexes[int(input[i + 1]) - 1]
    else:
        indexes[int(input[i]) - 1].next = indexes[len(input)]
indexes[-1].next = indexes[int(input[0]) - 1]

currentNumber = int(input[0])
for i in range(0, 10000000):
    next = indexes[currentNumber - 1].next
    indexes[currentNumber - 1].next = next.next.next.next

    destinationNumber = currentNumber - 1
    if destinationNumber == 0:
        destinationNumber = len(indexes)
    while destinationNumber == next.number or destinationNumber == next.next.number or destinationNumber == next.next.next.number:
        destinationNumber -= 1
        if destinationNumber == 0:
            destinationNumber = len(indexes)

    next.next.next.next = indexes[destinationNumber - 1].next
    indexes[destinationNumber - 1].next = next

    currentNumber = indexes[currentNumber - 1].next.number

print(indexes[0].next.number * indexes[0].next.next.number)