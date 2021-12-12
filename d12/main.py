from collections import defaultdict


def find_routes(tunnels: dict[str, set[str]], current: str = "start", visited: list[str] = ["start"], extra: bool = False) -> list[list[str]]:
    routes = []
    visited = visited.copy()

    for path in tunnels[current]:
        if path == "end":
            visited_new = visited.copy()
            visited_new.append("end")
            routes.append(visited_new)
        elif path == "start":
            continue
        elif path.islower() and path in visited:
            if path in visited and extra:
                continue
            else:
                visited_new = visited.copy()
                visited_new.append(path)
                routes += find_routes(tunnels, path, visited_new, True)
        else:
            visited_new = visited.copy()
            visited_new.append(path)
            routes += find_routes(tunnels, path, visited_new, extra)

    return routes


def get_tunnels() -> dict[str, set[str]]:
    t: dict[str, set[str]] = dict()
    with open("input.txt", "r") as f:
        for line in f.read().splitlines():
            a, b = line.split("-")
            if a in t:
                t[a].add(b)
            else:
                t[a] = set([b])
            if b in t:
                t[b].add(a)
            else:
                t[b] = set([a])
    return t


if __name__ == "__main__":
    tunnels = get_tunnels()
    print(tunnels)
    routes = find_routes(tunnels)
    print(len(routes))
