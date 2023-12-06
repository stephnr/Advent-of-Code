from functools import reduce

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def solve(p):
  return 0

p = load("input.txt")
print(f'Solution: {solve(p)}')