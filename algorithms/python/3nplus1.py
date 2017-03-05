import sys

buffer = []
for line in sys.stdin:
    line = line[:-1]
    buffer.append(line)

for line in buffer:
    numbers = line.split(' ')
    i = int(numbers[0])
    j = int(numbers[1])

    maxCount = 0
    for k in range(i,j):
        counter = 1
        while k > 1:
            if k % 2 == 0:
                k = k / 2
            else:
                k = 3 * k + 1
            counter += 1
        if counter > maxCount:
            maxCount = counter
    print(line + " " + str(maxCount))
