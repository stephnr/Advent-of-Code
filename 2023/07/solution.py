from functools import reduce
import functools

cardByScore = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A']

class Hand():
  def __init__(self, hand, bet) -> None:
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

  def cmpHigh(self, y) -> int:
    if self.strength != y.strength:
      return 0

    if self.high > y.high:
      return 1
    elif y.high > self.high:
      return -1

    for i in range(1, len(self.hand)):
      xi, yi = cardByScore.index(self.hand[i]), cardByScore.index(y.hand[i])

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

def solve(p):
  sum = 0
  hands = sorted([Hand(*r.split(' ')) for r in p], key=lambda x: x.strength)

  # Sort hands again comparing high cards of similar strengths
  hands = sorted(hands, key=functools.cmp_to_key(compareHandHighs))

  for i in range(0, len(hands)):
    hand = hands[i]
    # print(''.join(hand.hand), hand.strength, hand.high, hand.bet)
    sum += ((i+1) * int(hand.bet))

  return sum

p = load("input.txt")
print(f'Solution: {solve(p)}')