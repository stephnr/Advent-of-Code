def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def solve(p) -> int:
  return 0

print(f'Solution: ({solve(load("input.txt"))})')