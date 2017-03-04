import sys

buffer = []
for line in sys.stdin:
    line = line[:-1]
    buffer.append(line)

cases = buffer.pop(0)
current_index = 1
while current_index < len(buffer):
    candidates_count = int(buffer[current_index])
    current_index += 1
    candidates = buffer[current_index:current_index + candidates_count]
    current_index += candidates_count
    votes = []
    while buffer[current_index] != '':
        ballot = buffer[current_index].split()
        votes.append(ballot)
        current_index += 1
    tally = [0]*len(votes)
    for vote in votes:
        tally[int(vote[0])-1] += 1
    winner = ''
    excluded_candidates = []
    while winner == '':
        for index, candidate in tally:
            if candidate > len(votes/2):
                winner = candidates[index]

    print winner
    current_index += 30
print buffer
print cases
