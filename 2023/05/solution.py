def load(file):
  with open(file) as f:
    return [row.strip().split(':')[1] for row in f]


def solve(p):
  return 0


print(f'Solution: {solve(load("input.txt"))}')