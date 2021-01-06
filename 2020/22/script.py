file = open('input.txt','r')
lines = list(map(lambda line : line.strip(), file.readlines()))
players = [[],[]]

def createkey(p, players):
    key = ''
    for v in players[p]:
        key += "{:02d}".format(v)
    return key

player = players[0]
for line in lines:
    if line.isnumeric():
        player.append(int(line))
    elif line == '':
        player = players[1]

def game(players):

    playerDecks = [[], []]

    while len(players[0]) != 0 and len(players[1]) != 0:
        key = createkey(0, players)
        if key in playerDecks[0]:
            return 0
        playerDecks[0].append(key)
        key = createkey(1, players)
        if key in playerDecks[1]:
            return 0
        playerDecks[1].append(key)

        if len(players[0]) - 1 >= players[0][0] and len(players[1]) - 1 >= players[1][0]:
            winner = game([players[0][1:players[0][0] + 1],players[1][1:players[1][0] + 1]])
        else:
            if players[0][0] > players[1][0]:
                winner = 0
            else:
                winner = 1
        players[winner].append(players[winner].pop(0))
        players[winner].append(players[winner - 1].pop(0))

    if len(players[0]) > 0:
        return 0
    else:
        return 1

winner = players[game(players)]
score = 0
for i in range(0, len(winner)):
    score += winner[i] * (len(winner) - i)
print(score)