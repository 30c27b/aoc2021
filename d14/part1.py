from typing import Counter


def parse_input(path: str = "input.x=txt") -> tuple[str, dict[str, str]]:
    template: str = str()
    insertions: dict[str, str] = dict()
    with open(path, "r") as f:
        lines = f.read().splitlines()
        template = lines[0]
        for line in lines[2:]:
            k, v = line.split(" -> ")
            insertions[k] = v

    return template, insertions


def apply_step(template: str, insertions: dict[str, str]) -> str:
    new_template: str = str()
    for i, c in enumerate(template):
        new_template += c
        if i < len(template) - 1:
            new_template += insertions[c + template[i+1]]

    return new_template


if __name__ == "__main__":
    template, insertions = parse_input("input.txt")
    for i in range(40):
        template = apply_step(template, insertions)
    count = Counter(template)
    print(count.most_common(1)[0][1] - count.most_common()[-1][1])
