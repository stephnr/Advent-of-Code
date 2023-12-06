from functools import reduce

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def solve(p):
  times  = [int(el) for el in p[0].split(':')[1].strip().split(' ') if len(el) > 0]
  dists  = [int(el) for el in p[1].split(':')[1].strip().split(' ') if len(el) > 0]
  p2time = int(reduce(lambda x, y: ''+str(x)+str(y), times))
  p2dist = int(reduce(lambda x, y: ''+str(x)+str(y), dists))

  winPaths, p2winPath  = [0] * len(times), 0

  # Solve Part 1
  for i in range(0, len(times)):
    for speed in range(1, times[i]):
      if speed * (times[i] - speed) > dists[i]:
        winPaths[i] += 1

  # Solve Part 2
  for speed in range(1, p2time):
    if speed * (p2time - speed) > p2dist:
      p2winPath += 1

  return reduce(lambda x, y: x*y, winPaths), p2winPath

p = load("input.txt")
print(f'Solution: {solve(p)}')