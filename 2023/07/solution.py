from functools import reduce
import functools

cardByScore      = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A']
cardByScorePart2 = ['J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A']

class Hand():
  def __init__(self, part2, hand, bet) -> None:
    self.part2 = part2
    self.hand = [*hand]
    self.bet = bet

    # Find highest card
    self.high = cardByScore.index(self.hand[0])

    # Calculate hand strength
    uniq = list(set([*self.hand]))

    match len(uniq):
      case 1:
        self.strength = 6 # 5 of a kind
        pass

      case 2:
        self.strength = 4 # Full House

        for c in uniq:
          if self.hand.count(c) == 4:
            self.strength = 5 # Four Kind

        pass

      case 3:
        self.strength = 2 # Two Pair

        for c in uniq:
          if self.hand.count(c) == 3:
            self.strength = 3 # Three Kind

        pass

      case 4:
        self.strength = 1
        pass

      case 5:
        self.strength = 0
        pass

    if part2 and 'J' in self.hand:
      # Determine card rankings with Js
      match self.hand.count('J'):
        case 1:
          match len(uniq):
            case 2:
              self.strength = 6
              pass
            case 3:
              self.strength = 5
              pass
            case 4:
              self.strength = 4
              pass
            case 5:
              self.strength = 3
              pass
          pass

        case 2:
          match len(uniq):
            case 2:
              self.strength = 6
              pass
            case 3:
              self.strength = 5
              pass
            case 4:
              self.strength = 4
              pass
          pass

        case 3:
          if len(uniq) == 2:
            self.strength = 6
          else:
            self.strength = 5
          pass

        case 4:
          self.strength = 6
          pass

        case 5:
          self.strength = 0
          pass


  def cmpHigh(self, y) -> int:
    scoreTable = cardByScore

    if self.part2:
      scoreTable = cardByScorePart2

    # ---

    if self.strength != y.strength:
      return 0

    if self.high > y.high:
      return 1
    elif y.high > self.high:
      return -1

    for i in range(1, len(self.hand)):
      if self.part2:
        if self.hand.count('J') < y.hand.count('J'):
          return 1
        elif self.hand.count('J') > y.hand.count('J'):
          return -1

        if 'J' in self.hand and 'J' in y.hand:
          # If equal amounts of jacks - compare on non-jacks
          if len(list(set([*self.hand]))) == 2 and len(list(set([*y.hand]))) == 2:
            xi, yi = scoreTable.index(sorted(list(set([*self.hand])))[0]), scoreTable.index(sorted(list(set([*y.hand])))[0])
            return xi < yi

      xi, yi = scoreTable.index(self.hand[i]), scoreTable.index(y.hand[i])

      if xi > yi:
        return 1
      elif xi < yi:
        return -1

    return 0

def load(file):
  with open(file) as f:
    return [row.strip() for row in f]

def compareHandHighs(x: Hand, y: Hand) -> int:
  return x.cmpHigh(y)

def part1(p):
  sum = 0
  hands = sorted([Hand(False, *r.split(' ')) for r in p], key=lambda x: x.strength)

  # Sort hands again comparing high cards of similar strengths
  hands = sorted(hands, key=functools.cmp_to_key(compareHandHighs))

  for i in range(0, len(hands)):
    hand = hands[i]
    # print(''.join(hand.hand), hand.strength, hand.high, hand.bet)
    sum += ((i+1) * int(hand.bet))

  return sum

def part2(p):
  sum = 0
  hands = sorted([Hand(True, *r.split(' ')) for r in p], key=lambda x: x.strength)

  # Sort hands again comparing high cards of similar strengths
  hands = sorted(hands, key=functools.cmp_to_key(compareHandHighs))

  for i in range(0, len(hands)):
    hand = hands[i]
    print(''.join(hand.hand), hand.strength, hand.high, hand.bet)
    sum += ((i+1) * int(hand.bet))

  return sum

p = load("input.txt")
print(f'Solution: {part1(p), part2(p)}')