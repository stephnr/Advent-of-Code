# If source < destination before range add -> don't check (seed = destination)

import sys


class SeedMap():
  def __init__(self, dest, source, range) -> None:
    self.d = int(dest)
    self.s = int(source)
    self.r = int(range)

  def transform(self, seed) -> int:
    if seed >= self.s and seed <= self.s + self.r:
      return self.d + abs(seed - self.s)

    return seed

  def __str__(self) -> str:
    return f'Dest: {self.d}, Source: {self.s}, Range: {self.r}'

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]


def solve(p):
  # Parse Seeds
  seeds = [int(s) for s in (p[0].split(":")[1]).split()]
  # Store Seed Maps in order
  seedMaps = []

  seedMapParse = []

  for row in p[1:]:
    if 'map' in row:
      # Reset for new nested map
      if seedMapParse != []:
        seedMaps.append(seedMapParse)

      seedMapParse = []
      continue

    # Begin parsing
    if len(row) > 0:
      v = row.split()
      seedMapParse.append(SeedMap(*v))

  seedMaps.append(seedMapParse)

  # Seeds fully loaded
  lowestSeed = float('inf')

  for seed in seeds:
    print(f'\n\n>>NEW SEED: {seed}')
    for mapList in seedMaps:
      print("\n\n")
      for seedMap in mapList:
        print(f'Transforming - S: {seed} :: Map({seedMap})"')
        newSeed = seedMap.transform(seed)
        if newSeed != seed:
          print(f'{seed} -> {newSeed}')
          seed = newSeed
          break
    lowestSeed = min(seed, lowestSeed)

  return lowestSeed


print(f'Solution: {solve(load("input.txt"))}')