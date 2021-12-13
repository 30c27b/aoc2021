def parse_input(path: str = "input.x=txt") -> tuple[list[tuple[int, int]], list[tuple[bool, int]]]:
    points: list[tuple[int, int]] = []
    folds: list[tuple[bool, int]] = []
    with open(path, "r") as f:
        part2 = False
        for line in f.read().splitlines():
            if part2:
                a, b = line.split("=")
                folds.append((a.endswith("x"), int(b)))
            else:
                if line == "":
                    part2 = True
                    continue
                x, y = line.split(",")
                points.append((int(x), int(y)))

    return points, folds


def apply_fold(points: list[tuple[int, int]], fold: tuple[bool, int]):
    axis, n = fold
    # for x, y in points:
    #     new: tuple[int, int] = None
    #     if axis and x > n:
    #         new = (n-(x-n), y)
    #     elif not axis and y >= n:
    #         new = (n-(x-n), y)
    #     if new != None and new not in points:
    #         points.append(new)

    points = [point for point in points if point[0] == 1]


def display_points(points: list[tuple[int, int]], fold: tuple[bool, int] = None):
    mx = max(points, key=lambda t: t[0])[0]
    my = max(points, key=lambda t: t[1])[1]
    for y in range(my+1):
        for x in range(mx+1):
            if fold and not fold[0] and fold[1] == y:
                print("-", end="")
            elif fold and fold[0] and fold[1] == x:
                print("|", end="")
            else:
                print("#" if (x, y) in points else ".", end="")
        print()


if __name__ == "__main__":
    points, folds = parse_input("example_input.txt")
    display_points(points, folds[0])
    print("+++++++++++++++++")
    apply_fold(points, folds[0])
    display_points(points, folds[0])
