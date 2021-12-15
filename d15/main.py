from queue import PriorityQueue
import os
import sys

def parse_input(path: str = "input.txt") -> list[list[int]]:
    matrix:  list[list[int]] = list()
    with open(path, "r") as f:
        for line in f.read().splitlines():
            matrix.append([int(c) for c in line])

    return matrix


def init_queue(matrix: list[list[int]]):

    visited: list[tuple[int, int]] = []
    risk: list[list[int]] = [[-1 for x in range(len(line))] for line in matrix]

    visit((0, 0), matrix[0][0], risk, visited, matrix)
    
    print(risk)


def visit(pos: tuple[int, int], current_risk: int, risk: list[list[int]], visited: list[tuple[int, int]], matrix: list[list[int]]):
    if pos in visited:
        return
    x, y = pos
    if x < 0 or x >= len(matrix) or y < 0 or y >= len(matrix[0]):
        return
    print(x, y)
    r = current_risk + matrix[x][y]
    if risk[x][y] >= 0:
        risk[x][y] = min(r, risk[x][y])
    else:
        risk[x][y] = r
    visited.append(pos)

    visit((x, y-1), r, risk, visited, matrix)  # top
    visit((x+1, y), r, risk, visited, matrix)  # right
    visit((x, y+1), r, risk, visited, matrix)  # bot
    visit((x-1, y), r, risk, visited, matrix)  # left


if __name__ == "__main__":
    matrix = parse_input("input.txt")
    init_queue(matrix)
