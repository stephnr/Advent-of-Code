# If source < destination before range add -> don't check (seed = destination)

import sys

seeds, seedRanges, seedMaps = [], [], []
transformTable = {}

class SeedRange():
  def __init__(self, start, r) -> None:
    self.s = int(start)
    self.r = int(r)

class SeedMap():
  def __init__(self, dest, source, range) -> None:
    self.d = int(dest)
    self.s = int(source)
    self.r = int(range)

  def transformSeed(self, seed) -> int:
    if seed >= self.s and seed <= self.s + self.r:
      return self.d + seed - self.s

    return seed

  def inRange(self, sr: SeedRange) -> bool:
    return sr.s >= self.s and sr.s+sr.r <= self.s + self.r


  def __str__(self) -> str:
    return f'Dest: {self.d}, Source: {self.s}, Range: {self.r}'

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def parseMaps(p):
  global seeds
  global seedMaps
  global seedRange

  # Parse Seeds
  seeds = [int(s) for s in (p[0].split(":")[1]).split()]

  stmp = 0
  for seed in seeds:
    if stmp > 0:
      seedRanges.append(SeedRange(stmp, seed))
      stmp = 0
      continue

    stmp = max(stmp, seed)

  # Store Seed Maps in order
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

def part1():
  global transformTable

  lowestSeed = float('inf')

  for seed in seeds:
    initialSeed = seed
    # print(f'\n\n>>NEW SEED: {seed}')
    for mapList in seedMaps:
      # print("\n\n")
      for seedMap in mapList:
        # print(f'Transforming - S: {seed} :: Map({seedMap})"')
        newSeed = seedMap.transformSeed(seed)
        if newSeed != seed:
          # print(f'{seed} -> {newSeed}')
          seed = newSeed
          break

    transformTable[initialSeed] = seed
    lowestSeed = min(seed, lowestSeed)

  return lowestSeed

def part2():
  return 0 # can't seem to wrap my head around it v__v

  # global seedRanges

  # for mapList in seedMaps:
  #   newRanges = []

  #   for sm in mapList:
  #     for s in seedRanges:
  #       # With each seed map, compute a new seed range
  #       if sm.inRange(s):
  #         newStart = max(s.s, sm.s)
  #         newRange = min(s.s + s.r, sm.s + sm.r) - newStart
  #         newRanges.append(SeedRange(sm.transformSeed(newStart), newRange))
  #         continue
  #       else:
  #         # No match, so it corresponds to itself
  #         newRanges.append(s)

  #   seedRanges = newRanges




parseMaps(load("input.txt"))
print(f'Solution: ({part1()}, {part2()})')