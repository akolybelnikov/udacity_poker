package poker

import "strconv"

const handrankName = "HighCardPairTwoPairThreeOfAKindStraightFlushFullHouseFourOfAKindStraightFlush"

var handrankIndex = [...]uint8{0, 8, 12, 19, 31, 39, 44, 53, 64, 77}

func (i HandRank) String() string {
	if i < 0 || i >= HandRank(len(handrankIndex)-1) {
		return "HandRank(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return handrankName[handrankIndex[i]:handrankIndex[i+1]]
}

const rankName = "TwoThreeFourFiveSixSevenEightNineTenJackQueenKingAce"

var rankIndex = [...]uint8{0, 3, 8, 12, 16, 19, 24, 29, 33, 36, 40, 45, 49, 52}

func (i Rank) String() string {
	i -= 2
	if i < 0 || i >= Rank(len(rankIndex)-1) {
		return "Rank(" + strconv.FormatInt(int64(i+2), 10) + ")"
	}
	return rankName[rankIndex[i]:rankIndex[i+1]]
}

const suitName = "SpadeDiamondClubHeart"

var suitIndex = [...]uint8{0, 5, 12, 16, 21}

func (i Suit) String() string {
	if i < 0 || i >= Suit(len(suitIndex)-1) {
		return "Suit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return suitName[suitIndex[i]:suitIndex[i+1]]
}
