package poker

import (
	"fmt"
	"math/rand"
)

func deck() []Card {
	cards := make([]Card, 52)
	for i := 0; i < 52; i++ {
		cards[i].rank = Rank(i%13 + 2)
		cards[i].suit = Suit(i / 13)
	}
	return cards
}

func deal(numHands int) ([]Hand, error) {
	cards := deck()
	if 5*numHands > len(cards) {
		return nil, fmt.Errorf("not enough cards in the deck")
	}
	// shuffle the deck
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	hands := make([]Hand, numHands)
	for i := range hands {
		hands[i].cards = [5]Card(cards[i*5 : (i+1)*5])
	}

	return hands, nil
}

func play(hands []Hand) []Hand {
	var max Hand
	winners := make([]Hand, 0)

	for _, h := range hands {
		h.score()
		if len(winners) == 0 {
			winners = append(winners, h)
			max = h
			continue
		}
		if h.handRank > max.handRank {
			winners = winners[:0]
			winners = append(winners, h)
			max = h
			continue
		}
		if h.handRank == max.handRank {
			var idx int
			for i, r := range h.ranks {
				if r > max.ranks[i] {
					winners = winners[:0]
					winners = append(winners, h)
					max = h
					break
				}
				if r < max.ranks[i] {
					break
				}
				idx++
			}
			if idx == 5 {
				winners = append(winners, h)
			}
		}
	}

	return winners
}

func poker(numHands int) error {
	hands, err := deal(numHands)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	for _, h := range hands {
		fmt.Printf("%s\n", h.String())
	}

	winners := play(hands)
	fmt.Printf("%d Winner(s):\n", len(winners))

	for _, w := range winners {
		fmt.Printf("%s: %s, %s\n", w.String(), w.handRank, w.ranks)
	}

	return nil
}
