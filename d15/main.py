from queue import PriorityQueue


def parse_input(path: str = "input.txt") -> list[list[int]]:
    risk:  list[list[int]] = list()
    with open(path, "r") as f:
        for line in f.read().splitlines():
            risk.append([int(c) for c in line])

    return risk


def part2(risk: list[list[int]]) -> list[list[int]]:
    new_risk: list[list[int]] = []

    for x in range(len(risk)*5):
        new_risk.append([0 for y in range(len(risk)*5)])

    for x in range(len(new_risk)):
        for y in range(len(new_risk)):
            new_risk[x][y] = risk[x % len(risk)][y % len(
                risk)] + (x // len(risk)) + (y // len(risk))
            if new_risk[x][y] > 9:
                new_risk[x][y] = (new_risk[x][y] % 10) + 1

    return new_risk


def set_totalrisk(current: int, point: tuple[int, int], closedl: list[list[int]], risk: list[list[int]]) -> bool:
    x, y = point
    if x < 0 or x >= len(risk) or y < 0 or y >= len(risk[0]):
        return
    r = current + risk[x][y]

    if closedl[x][y] == -1:
        closedl[x][y] = r
    elif r < closedl[x][y]:
        closedl[x][y] = r
    else:
        return False
    return True


def astar(risk: list[list[int]]):
    openl = PriorityQueue()
    closedl: list[list[int]] = [
        [-1 for x in range(len(line))] for line in risk]

    openl.put((0, 0), 0)
    closedl[0][0] = 0

    while not openl.empty():

        point = openl.get()
        x, y = point
        if point == (len(risk) - 1, len(risk) - 1):
            print(closedl[x][y])
            return

        if set_totalrisk(closedl[x][y], (x, y-1), closedl, risk):
            openl.put((x, y-1), closedl[x][y-1] +
                      (len(risk) - x) + (len(risk)-(y-1)))
        if set_totalrisk(closedl[x][y], (x+1, y), closedl, risk):
            openl.put((x+1, y), closedl[x+1][y] +
                      (len(risk) - (x + 1)) + (len(risk)-y))
        if set_totalrisk(closedl[x][y], (x-1, y), closedl, risk):
            openl.put((x-1, y), closedl[x-1][y] +
                      (len(risk) - (x-1)) + (len(risk)-y))
        if set_totalrisk(closedl[x][y], (x, y+1), closedl, risk):
            openl.put((x, y+1), closedl[x][y+1] +
                      (len(risk) - x) + (len(risk)-(y+1)))


if __name__ == "__main__":
    risk = parse_input("input.txt")
    risk = part2(risk)
    astar(risk)
