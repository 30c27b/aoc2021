def main():
    f = open("input.txt", "r")
    lines = f.readlines()
    lines = [line.rstrip() for line in lines]
    f.close()

    grid = []

    for line in lines:
        grid.append([(int(c), False) for c in line])

    stepno = 1
    while True:
        grid, f = step(grid)
        if f == 100:
            break
        stepno += 1

    print(stepno)


def step(grid: list[tuple[int, bool]]) -> list[list[tuple[int, bool], int]]:
    flashes = 0

    for x in range(len(grid)):
        grid[x] = [(n[0] + 1, False) for n in grid[x]]

    for x in range(len(grid)):
        for y in range(len(grid[0])):
            if grid[x][y][0] >= 10 and not grid[x][y][1]:
                grid, f = flash(grid, (x, y))
                flashes += f

    return grid, flashes


def flash(grid: list[tuple[int, bool]], el: tuple[int, int]) -> list[list[tuple[int, bool], int]]:
    x, y = el
    flashes = 0

    grid[x][y] = (0, True)

    if x - 1 >= 0 and not grid[x-1][y][1]:
        grid[x-1][y] = (grid[x-1][y][0] + 1, grid[x-1][y][1])
        if grid[x-1][y][0] >= 10:
            grid, f = flash(grid, (x-1, y))
            flashes += f

    if x + 1 < len(grid) and not grid[x+1][y][1]:
        grid[x+1][y] = (grid[x+1][y][0] + 1, grid[x+1][y][1])
        if grid[x+1][y][0] >= 10:
            grid, f = flash(grid, (x+1, y))
            flashes += f

    if y - 1 >= 0 and not grid[x][y-1][1]:
        grid[x][y-1] = (grid[x][y-1][0] + 1, grid[x][y-1][1])
        if grid[x][y-1][0] >= 10:
            grid, f = flash(grid, (x, y-1))
            flashes += f

    if y + 1 < len(grid[0]) and not grid[x][y+1][1]:
        grid[x][y+1] = (grid[x][y+1][0] + 1, grid[x][y+1][1])
        if grid[x][y+1][0] >= 10:
            grid, f = flash(grid, (x, y+1))
            flashes += f

    if x - 1 >= 0 and y - 1 >= 0 and not grid[x-1][y-1][1]:
        grid[x-1][y-1] = (grid[x-1][y-1][0] + 1, grid[x-1][y-1][1])
        if grid[x-1][y-1][0] >= 10:
            grid, f = flash(grid, (x-1, y-1))
            flashes += f

    if x + 1 < len(grid) and y - 1 >= 0 and not grid[x+1][y-1][1]:
        grid[x+1][y-1] = (grid[x+1][y-1][0] + 1, grid[x+1][y-1][1])
        if grid[x+1][y-1][0] >= 10:
            grid, f = flash(grid, (x+1, y-1))
            flashes += f

    if x - 1 >= 0 and y + 1 < len(grid[0]) and not grid[x-1][y+1][1]:
        grid[x-1][y+1] = (grid[x-1][y+1][0] + 1, grid[x-1][y+1][1])
        if grid[x-1][y+1][0] >= 10:
            grid, f = flash(grid, (x-1, y+1))
            flashes += f

    if x + 1 < len(grid) and y + 1 < len(grid[0]) and not grid[x+1][y+1][1]:
        grid[x+1][y+1] = (grid[x+1][y+1][0] + 1, grid[x+1][y+1][1])
        if grid[x+1][y+1][0] >= 10:
            grid, f = flash(grid, (x+1, y+1))
            flashes += f

    return grid, flashes + 1


main()
