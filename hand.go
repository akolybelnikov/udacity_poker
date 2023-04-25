package poker

import (
	"fmt"
	"sort"
)

type Hand struct {
	cards    [5]Card
	count    map[Rank]int
	handRank HandRank
	ranks    [5]Rank
}

func (h *Hand) flush() bool {
	suit := h.cards[0].suit
	for i, c := range h.cards {
		if c.suit != suit {
			h.ranks = [5]Rank{}
			return false
		}
		h.ranks[i] = c.rank
	}
	h.handRank = Flush

	return true
}

func (h *Hand) straight() bool {
	rank := h.cards[0].rank
	if rank == 14 && h.cards[1].rank == 5 {
		h.ranks = [5]Rank{5, 4, 3, 2, 1}
		h.handRank = Straight
		return true
	}
	for i, c := range h.cards {
		if c.rank != rank {
			return false
		}
		h.ranks[i] = c.rank
		rank--
	}
	h.handRank = Straight

	return true
}

func (h *Hand) kind() HandRank {
	h.count = make(map[Rank]int)
	for _, c := range h.cards {
		h.count[c.rank]++
	}

	switch len(h.count) {
	case 2:
		return h.rankTwo()
	case 3:
		return h.rankThree()
	case 4:
		return h.rankFour()
	default:
		return h.rankHighCard()
	}
}

func (h *Hand) rankTwo() HandRank {
	var r HandRank
	for k, v := range h.count {
		switch v {
		case 1:
			r = FourOfAKind
			h.ranks[1] = k
		case 2:
			r = FullHouse
			h.ranks[1] = k
		case 3:
			r = FullHouse
			h.ranks[0] = k
		case 4:
			r = FourOfAKind
			h.ranks[0] = k
		}
	}

	return r
}

func (h *Hand) rankThree() HandRank {
	var r HandRank
	for _, v := range h.count {
		switch v {
		case 2:
			r = TwoPair
			h.rankTwoPair()
		case 3:
			r = ThreeOfAKind
			h.rankThreeOfAKind()
		}
	}

	return r
}

func (h *Hand) rankFour() HandRank {
	rs := make([]Rank, 0)
	for k, v := range h.count {
		switch v {
		case 2:
			h.ranks[0] = k
		case 1:
			rs = append(rs, k)
		}
	}
	sort.Slice(rs, func(i, j int) bool { return rs[i] > rs[j] })
	for i, r := range rs {
		h.ranks[i+1] = r
	}

	return Pair
}

func (h *Hand) rankThreeOfAKind() {
	for k, v := range h.count {
		switch v {
		case 1:
			if h.ranks[1] == 0 {
				h.ranks[1] = k
			} else if h.ranks[1] < k {
				h.ranks[2] = h.ranks[1]
				h.ranks[1] = k
			} else {
				h.ranks[2] = k
			}
		case 3:
			h.ranks[0] = k
		}
	}
}

func (h *Hand) rankTwoPair() {
	var r1, r2 Rank
	for k, v := range h.count {
		switch v {
		case 1:
			h.ranks[2] = k
		case 2:
			if r1 == 0 {
				r1 = k
			} else {
				if r1 < k {
					r2 = r1
					r1 = k
				} else {
					r2 = k
				}
			}
		}
	}
	h.ranks[0] = r1
	h.ranks[1] = r2
}

func (h *Hand) rankHighCard() HandRank {
	for i, card := range h.cards {
		h.ranks[i] = card.rank
	}

	return HighCard
}

func (h *Hand) score() {
	h.sort()
	f := h.flush()
	s := h.straight()
	switch {
	case f && s:
		h.handRank = StraightFlush
	case f:
		h.handRank = Flush
	case s:
		h.handRank = Straight
	default:
		h.handRank = h.kind()
	}
}

func (h *Hand) sort() {
	sort.SliceStable(h.cards[:], func(i, j int) bool {
		return h.cards[i].rank > h.cards[j].rank
	})
}

func (h *Hand) String() string {
	buf := ""
	for _, c := range h.cards {
		buf += fmt.Sprintf("%s:%s ", c.rank, c.suit)
	}

	return buf
}
