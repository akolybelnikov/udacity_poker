package poker

import (
	"sort"
	"testing"
)

func TestHand_flush(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]Card
		want  bool
	}{
		{
			name: "flush of spades",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Two, suit: Spade},
				{rank: Queen, suit: Spade},
				{rank: Five, suit: Spade},
				{rank: Ten, suit: Spade},
			},
			want: true,
		},
		{
			name: "no flush",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Two, suit: Spade},
				{rank: Queen, suit: Spade},
				{rank: Five, suit: Spade},
				{rank: Ten, suit: Heart},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.cards,
			}
			if got := h.flush(); got != tt.want {
				t.Errorf("flush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_kind(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]Card
		want  HandRank
	}{
		{
			name: "four of a kind",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Ace, suit: Diamond},
				{rank: Ten, suit: Spade},
			},
			want: FourOfAKind,
		},
		{
			name: "full house",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Ten, suit: Spade},
			},
			want: FullHouse,
		},
		{
			name: "three of a kind",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Two, suit: Spade},
			},
			want: ThreeOfAKind,
		},
		{
			name: "two pair",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Two, suit: Club},
				{rank: Two, suit: Spade},
			},
			want: TwoPair,
		},
		{
			name: "pair",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Two, suit: Club},
				{rank: Three, suit: Spade},
			},
			want: Pair,
		},
		{
			name: "high card",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: King, suit: Heart},
				{rank: Two, suit: Club},
				{rank: Three, suit: Spade},
			},
			want: HighCard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.cards,
			}
			if got := h.kind(); got != tt.want {
				t.Errorf("kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_rankThree(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]Card
		want  [5]Rank
	}{
		{
			name: "three of a kind",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Two, suit: Diamond},
				{rank: Ten, suit: Spade},
			},
			want: [5]Rank{Ace, Ten, Two, Rank(0), Rank(0)},
		},
		{
			name: "two pair",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Two, suit: Club},
				{rank: Two, suit: Diamond},
				{rank: Ten, suit: Spade},
			},
			want: [5]Rank{Ace, Two, Ten, Rank(0), Rank(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.cards,
			}
			h.score()
			if got := h.ranks; got != tt.want {
				t.Errorf("rankThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_rankTwo(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]Card
		want  [5]Rank
	}{
		{
			name: "four of a kind",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Ace, suit: Diamond},
				{rank: Ten, suit: Spade},
			},
			want: [5]Rank{Ace, Ten, Rank(0), Rank(0), Rank(0)},
		},
		{
			name: "full house",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Ten, suit: Spade},
			},
			want: [5]Rank{Ace, Ten, Rank(0), Rank(0), Rank(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.cards,
			}
			h.score()
			if got := h.ranks; got != tt.want {
				t.Errorf("rankTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_score(t *testing.T) {
	tests := []struct {
		name      string
		cards     [5]Card
		wantScore HandRank
		wantRanks [5]Rank
	}{
		{
			name: "straight flush",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Jack, suit: Diamond},
				{rank: Queen, suit: Diamond},
				{rank: King, suit: Diamond},
				{rank: Ace, suit: Diamond},
			},
			wantScore: StraightFlush,
			wantRanks: [5]Rank{Ace, King, Queen, Jack, Ten},
		},
		{
			name: "straight flush ace low",
			cards: [5]Card{
				{rank: Two, suit: Diamond},
				{rank: Five, suit: Diamond},
				{rank: Four, suit: Diamond},
				{rank: Three, suit: Diamond},
				{rank: Ace, suit: Diamond},
			},
			wantScore: StraightFlush,
			wantRanks: [5]Rank{Five, Four, Three, Two, Rank(1)},
		},
		{
			name: "four of a kind",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Ace, suit: Diamond},
				{rank: Ten, suit: Spade},
			},
			wantScore: FourOfAKind,
			wantRanks: [5]Rank{Ace, Ten, Rank(0), Rank(0), Rank(0)},
		},
		{
			name: "full house",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Ace, suit: Spade},
				{rank: Ace, suit: Heart},
				{rank: Ace, suit: Club},
				{rank: Ten, suit: Spade},
			},
			wantScore: FullHouse,
			wantRanks: [5]Rank{Ace, Ten, Rank(0), Rank(0), Rank(0)},
		},
		{
			name: "flush",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Two, suit: Diamond},
				{rank: Three, suit: Diamond},
				{rank: Four, suit: Diamond},
				{rank: Ace, suit: Diamond},
			},
			wantScore: Flush,
			wantRanks: [5]Rank{Ace, Ten, Four, Three, Two},
		},
		{
			name: "straight",
			cards: [5]Card{
				{rank: Ten, suit: Diamond},
				{rank: Jack, suit: Spade},
				{rank: Queen, suit: Heart},
				{rank: King, suit: Club},
				{rank: Ace, suit: Diamond},
			},
			wantScore: Straight,
			wantRanks: [5]Rank{Ace, King, Queen, Jack, Ten},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.cards,
			}
			h.score()
			if got := h.handRank; got != tt.wantScore {
				t.Errorf("score() = %v, want %v", got, tt.wantScore)
			}
			if got := h.ranks; got != tt.wantRanks {
				t.Errorf("score() = %v, want %v", got, tt.wantRanks)
			}
		})
	}
}

func TestHand_straight(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]Card
		want  bool
	}{
		{
			name: "no straight",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Two, suit: Diamond},
				{rank: Three, suit: Heart},
				{rank: Four, suit: Club},
				{rank: Six, suit: Spade},
			},
			want: false,
		},
		{
			name: "straight",
			cards: [5]Card{
				{rank: Six, suit: Spade},
				{rank: Two, suit: Diamond},
				{rank: Three, suit: Heart},
				{rank: Four, suit: Club},
				{rank: Five, suit: Spade},
			},
			want: true,
		},
		{
			name: "straight with ace low",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Two, suit: Diamond},
				{rank: Three, suit: Heart},
				{rank: Four, suit: Club},
				{rank: Five, suit: Spade},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.SliceStable(tt.cards[:], func(i, j int) bool {
				return tt.cards[i].rank > tt.cards[j].rank
			})
			h := Hand{
				cards: tt.cards,
			}
			if got := h.straight(); got != tt.want {
				t.Errorf("straight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_String(t *testing.T) {
	tests := []struct {
		name string
		i    Rank
		want string
	}{
		{
			name: "Two",
			i:    Two,
			want: "Two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuit_String(t *testing.T) {
	tests := []struct {
		name string
		i    Suit
		want string
	}{
		{
			name: "Diamond",
			i:    Diamond,
			want: "Diamond",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandRank_String(t *testing.T) {
	tests := []struct {
		name string
		i    HandRank
		want string
	}{
		{
			name: "Four of a kind",
			i:    FourOfAKind,
			want: "FourOfAKind",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_String(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]Card
		want  string
	}{
		{
			name: "hand 1",
			cards: [5]Card{
				{rank: Ace, suit: Spade},
				{rank: Two, suit: Diamond},
				{rank: Three, suit: Heart},
				{rank: Four, suit: Club},
				{rank: Five, suit: Spade},
			},
			want: "Ace:Spade Two:Diamond Three:Heart Four:Club Five:Spade ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.cards,
			}
			if got := h.String(); got != tt.want {
				t.Errorf("String() = %v with len %d, want %v with len %d", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}
