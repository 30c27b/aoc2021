
def main():
    f = open("input.txt", "r")
    lines = f.readlines()
    lines = [line.rstrip() for line in lines]
    f.close()
    height = len(lines)
    width = len(lines[0])

    lows = []

    for i in range(len(lines)):
        for j in range(len(lines[i])):
            isLow = True
            if i > 0 and lines[i - 1][j] <= lines[i][j]:
                isLow = False
            if i < height - 1 and lines[i + 1][j] <= lines[i][j]:
                isLow = False
            if j > 0 and lines[i][j - 1] <= lines[i][j]:
                isLow = False
            if j < width - 1 and lines[i][j + 1] <= lines[i][j]:
                isLow = False
            if isLow:
                lows.append((i, j))


main()
