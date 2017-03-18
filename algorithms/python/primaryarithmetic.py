def carry(twoNumbers):
    firstNumber = twoNumbers[0][::-1]
    secondNumber = twoNumbers[1][::-1]

    carrys = 0
    previousCarry = 0
    longerLength = len(firstNumber)
    if len(firstNumber) < len(secondNumber):
        longerLength = len(secondNumber)
    for i in range(0, longerLength - 1):
        first = 0
        second = 0
        if i < len(firstNumber):
            first = int(firstNumber[i])
        if i < len(secondNumber):
            second = int(secondNumber[i])

        if first + second + previousCarry > 9:
            previousCarry = 1
            carrys += 1
        else:
            previousCarry = 0

        if i >= len(firstNumber) or i >= len(secondNumber):
            if previousCarry == 1 and (first == 9 or second == 9):
                previousCarry = 1
                carrys += 1

    print(carrys)


carry(['123', '456'])
carry(['555', '555'])
carry(['123', '594'])
carry(['12338292', '594'])
