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
  return 0

p = load("input.txt")
print(f'Solution: {part1(p), part2(p)}')