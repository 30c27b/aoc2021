chunks = {
    '(': ')',
    '[': ']',
    '{': '}',
    '<': '>',
}

scores = {
    ')': 1,
    ']': 2,
    '}': 3,
    '>': 4,
}

def main():

    f = open("input.txt", "r")
    lines = f.readlines()
    lines = [line.rstrip() for line in lines]
    f.close()

    mcs = []

    for line in lines:

        i = 0
        legal = True
        mc = []
        while i < len(line) and legal:
            legal, mc, at = parse_chunk(line[i:])
            i += at + 1

        if legal:
            mcs.append(mc)
        
    totals = []

    for line in mcs:

        pass
        scrs = [scores[mc] for mc in line]
        score = 0
        for sc in scrs:
            score *= 5
            score += sc
        totals.append(score)

    totals = sorted(totals)

    print(totals[int(len(totals)/2)])



# return is [legal, list of missing chars, current char index]

def parse_chunk(line: str) -> tuple[bool, list[str], int]:

    openc = line[0]
    closec = chunks[openc]

    mcs = []

    i = 1

    while i < len(line):

        if line[i] == closec:  # if this chunk is closed

            return True, [], i

        elif line[i] in chunks:  # else if a new chunk starts inside the current chunk

            legal, mc, x = parse_chunk(line[i:])
            if not legal:
                return False, None, -1
            mcs += mc
            i += x

        else: # else, an illegal character has been found (should be a wrong closing statement)

            return False, None, -1

        i += 1

    # end of line with missing end of chunk
    mcs.append(closec)
    return True, mcs, i


main()
