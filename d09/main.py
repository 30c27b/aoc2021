from enum import Enum
import math


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

    bassins = []

    for x, y in lows:
        print("~~~~~~~~~~ bassin", x, y)
        a = create_marked_array(lines)
        bassins.append(bassin_size(a, (x, y)))
        print_marked_array(a)

    biggest = sorted(bassins, reverse=True)[:3]

    print(math.prod(biggest))


def create_marked_array(lines: list[str]) -> list[list[tuple[int, bool]]]:
    x_max = len(lines)
    y_max = len(lines[0])

    array = []

    for line in lines:
        array.append([(int(c), False) for c in line])

    return array

def print_marked_array(a: list[list[tuple[int, bool]]]):

    for l in a:
        for v, m in l:
            if m:
                print('\033[92m', v, '\033[0m' ,sep='', end='')
            else:
                print(v, end='')
            pass
        print()


def bassin_size(array: list[list[tuple[int, bool]]], point: tuple[int, int]):
    x_max = len(array)
    y_max = len(array[0])
    x, y = point

    if array[x][y][1]:
        return 0

    n = array[x][y][0]
    array[x][y] = (n, True)

    if n == 9:
        return 0

    size = 1

    if x > 0:
        size += bassin_size(array, (x - 1, y))

    if x < x_max - 1:
        size += bassin_size(array, (x + 1, y))

    if y > 0:
        size += bassin_size(array, (x, y - 1))

    if y < y_max - 1:
        size += bassin_size(array, (x, y + 1))

    return size


main()
