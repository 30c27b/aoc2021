from typing import Counter
from collections import defaultdict


def parse_input(path: str = "input.x=txt") -> tuple[str, dict[str, str]]:
    template: str = str()
    insertions: dict[str, str] = dict()
    with open(path, "r") as f:
        lines = f.read().splitlines()
        template = lines[0]
        for line in lines[2:]:
            k, v = line.split(" -> ")
            insertions[k] = v

    lastChar = template[-1]
    return template, insertions


def init_pairs(template: str, insertions: dict[str, str]) -> dict[str, int]:
    pairs: dict[str, int] = dict()
    for k, v in insertions.items():
        pairs[k] = 0

    for i, c in enumerate(template):
        if i >= len(template) - 1:
            break
        pairs[c + template[i+1]] += 1

    return pairs


def apply_step(pairs: dict[str, int], insertions: dict[str, str]) -> dict[str, int]:
    new_pairs: dict[str, int] = defaultdict(lambda: 0)
    for k, v in pairs.items():
        new_pairs[k[0] + insertions[k]] += v
        new_pairs[insertions[k] + k[1]] += v

    return new_pairs


def count_pairs(pairs: dict[str, int], insertions: dict[str, str], template: str) -> dict[str, str]:
    c: dict[str, int] = defaultdict(lambda: 0)
    for k,v in pairs.items():
        c[k[0]] += v
    c[template[-1]] += 1
    return c



if __name__ == "__main__":
    template, insertions = parse_input("input.txt")
    pairs: dict[str, int] = init_pairs(template, insertions)
    for _ in range(40):
        pairs = apply_step(pairs, insertions)
    c = count_pairs(pairs, insertions, template)
    count = Counter(c)
    print(count.most_common(1)[0][1] - count.most_common()[-1][1])
