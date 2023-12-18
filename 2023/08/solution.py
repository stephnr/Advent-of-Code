from functools import reduce
from math import gcd
import re

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def loadNodes(lines):
  nodes = {}

  for line in range(2, len(lines)):
    node, x, left, right = re.sub(r'\(?\)?\,?\=?', '', p[line]).split(' ')
    nodes[node] = tuple([left, right])

  return nodes

def lcm(a, b):
  return a * b // gcd(a, b)

def part1(p):
  steps = list(p[0])
  nodes = loadNodes(p)

  # Start traversal
  node, count = nodes['AAA'], 0

  while True:
    step = steps[count % len(steps)]
    count += 1

    match step:
      case 'L':
        if node[0] == 'ZZZ':
          break

        node = nodes[node[0]]
        pass
      case 'R':
        if node[1] == 'ZZZ':
          break

        node = nodes[node[1]]
        pass

  return count

def part2(p):
  steps = list(p[0])
  nodes = loadNodes(p)

  # Start traversal
  node, count = [node for node in nodes if node[-1] == "A"], 0
  steps_needed = []

  # Find steps needed to reach each of the Z nodes
  while len(node) > 0:
    step = steps[count % len(steps)]
    count += 1

    newNodes = []

    for n in node:
      match step:
        case 'L':
          if n[-1] == 'Z':
            steps_needed.append(count-1)
          else:
            newNodes.append(nodes[n][0])

          pass
        case 'R':
          if n[-1] == 'Z':
            steps_needed.append(count-1)
          else:
            newNodes.append(nodes[n][1])

          pass

    node = newNodes

  return reduce(lcm, steps_needed)

p = load("input.txt")
print(f'Solution: {part1(p), part2(p)}')