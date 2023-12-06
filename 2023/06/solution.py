from functools import reduce

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def solve(p) -> int:
  times = [int(el) for el in p[0].split(':')[1].strip().split(' ') if len(el) > 0]
  distances = [int(el) for el in p[1].split(':')[1].strip().split(' ') if len(el) > 0]
  winPaths = [0] * len(times)

  for i in range(0, len(times)):
    for speed in range(1, times[i]):
      remainingTime = times[i] - speed

      if speed * remainingTime > distances[i]:
        winPaths[i] += 1

  return reduce(lambda x, y: x*y, winPaths)

print(f'Solution: ({solve(load("input.txt"))})')