def load(file):
  with open(file) as f:
    return [row.strip().split(':')[1] for row in f]


def solve(p):
  # Form result trackers & scratch tables to track copies
  part1, cards = 0, [1] * len(p)

  for id, row in enumerate(p):
    # Count the number of winners by reading the length of an intersection
    if winner := len(set.intersection(*[set(map(int,part.split())) for part in row.split('|')])):
      # Part 1 => counting the score exponentially on base 2
      part1 += 2**(winner-1)
      # Part 2 => Count the number of cards formed without nesting by ranging over
      for n in range(1, winner+1):
        cards[id+n] += cards[id]
  return part1, sum(cards)


print(f'Solutions: {solve(load("input.txt"))}')