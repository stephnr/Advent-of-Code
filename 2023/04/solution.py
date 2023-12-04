def load(file):
  with open(file) as f:
    return [row.strip().split(':')[1] for row in f]


def solve(p):
  part1, cards = 0, [1] * len(p)
  for id, row in enumerate(p):
    # Count the number of winners by reading the length of an intersection
    if winner := len(set.intersection(*[set(map(int,part.split())) for part in row.split('|')])):
      # Part 1 => counting the score exponentially on base 2
      part1 += 2**(winner-1)
      # Part 2 => Form a dynamic range and count the number of cards formed without nesting
      for n in range(1, winner+1):
        cards[id+n] += cards[id]
  return part1, sum(cards)


print(f'Part 1: {solve(load("input.txt"))}')