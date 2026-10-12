package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "one":
        return 1
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten", "jack", "queen", "king":
        return 10
    case "ace":
        return 11
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Int := ParseCard(card1)
    card2Int := ParseCard(card2)
    dealerCardInt := ParseCard(dealerCard)
    sumOfPlayerCards := card1Int + card2Int
    switch {
    case card1Int == 11 && card2Int == 11:
        return "P"
    case sumOfPlayerCards == 21:
        if dealerCardInt < 10 {
            return "W"
        } else {
            return "S"
        }
    case sumOfPlayerCards >= 17 && sumOfPlayerCards <= 20:
        return "S"
    case sumOfPlayerCards >= 12 && sumOfPlayerCards <= 16:
        if dealerCardInt >= 7 {
            return "H"
        } else {
            return "S"
        }
    case sumOfPlayerCards <= 11:
        return "H"
    default:
        return ""
    }
}
