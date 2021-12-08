from typing import Mapping


def main():
    f = open("input.txt", "r")

    lines = f.readlines()
    lines = [line.rstrip() for line in lines]

    f.close()

    total = 0
    for line in lines:
        parts = line.split(" | ")
        total += mapSegments(parts[0].split(" "), parts[1].split(" "))

    print(total)


def mapSegments(inputs: list[str], codes: list[str]) -> int:

    digits = ["".join(sorted(x)) for x in inputs]
    # print(digits)

    numbers = {
        0: "",
        1: "",
        2: "",
        3: "",
        4: "",
        5: "",
        6: "",
        7: "",
        8: "",
        9: "",
    }

    two_tree_five = []
    zero_six_nine = []

    for digit in digits:
        if len(digit) == 2:
            numbers[1] = digit
        elif len(digit) == 3:
            numbers[7] = digit
        elif len(digit) == 4:
            numbers[4] = digit
        elif len(digit) == 5:
            two_tree_five += [digit]
        elif len(digit) == 6:
            zero_six_nine += [digit]
        elif len(digit) == 7:
            numbers[8] = digit

    for digit in two_tree_five:
        # 3 is the only case whre all chars of 1 are in digit
        if set(numbers[1] + digit) == set(digit):
            numbers[3] = digit
        # all segments for x + 4 only if x is 2
        elif set(digit + numbers[4]) == set(numbers[8]):
            numbers[2] = digit
        else:  # only remaining possibility
            numbers[5] = digit

    for digit in zero_six_nine:
        if set(numbers[1] + digit) == set(numbers[8]):  # x + 1 is full only for x=6
            numbers[6] = digit
        # x + 4 does not change only for 9
        elif set(digit + numbers[4]) == set(digit):
            numbers[9] = digit
        else:
            numbers[0] = digit

    rev = {value: key for key, value in numbers.items()}

    return int("".join([str(rev["".join(sorted(x))]) for x in codes]))



main()
