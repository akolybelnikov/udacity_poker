package poker

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_play(t *testing.T) {
	type args struct {
		hands []Hand
	}
	tests := []struct {
		name string
		args args
		want []Hand
	}{
		{
			name: "HighCard one winner",
			args: args{
				hands: []Hand{
					{
						cards: [5]Card{
							{rank: Ace, suit: Spade},
							{rank: Two, suit: Spade},
							{rank: Queen, suit: Spade},
							{rank: Five, suit: Spade},
							{rank: Ten, suit: Heart},
						},
					},
					{
						cards: [5]Card{
							{rank: King, suit: Heart},
							{rank: Two, suit: Diamond},
							{rank: Queen, suit: Club},
							{rank: Five, suit: Spade},
							{rank: Ten, suit: Club},
						},
					},
				},
			},
			want: []Hand{
				{
					cards: [5]Card{
						{rank: Ace, suit: Spade},
						{rank: Queen, suit: Spade},
						{rank: Ten, suit: Heart},
						{rank: Five, suit: Spade},
						{rank: Two, suit: Spade},
					},
					ranks: [5]Rank{
						Ace,
						Queen,
						Ten,
						Five,
						Two,
					},
					handRank: HighCard,
					count: map[Rank]int{
						Ace:   1,
						Queen: 1,
						Ten:   1,
						Five:  1,
						Two:   1,
					},
				},
			},
		},
		{
			name: "HighCard two winners",
			args: args{
				hands: []Hand{
					{
						cards: [5]Card{
							{rank: Ace, suit: Spade},
							{rank: Two, suit: Spade},
							{rank: Queen, suit: Spade},
							{rank: Five, suit: Spade},
							{rank: Ten, suit: Heart},
						},
					},
					{
						cards: [5]Card{
							{rank: Ace, suit: Heart},
							{rank: Two, suit: Diamond},
							{rank: Queen, suit: Club},
							{rank: Five, suit: Spade},
							{rank: Ten, suit: Club},
						},
					},
				},
			},
			want: []Hand{
				{
					cards: [5]Card{
						{rank: Ace, suit: Spade},
						{rank: Queen, suit: Spade},
						{rank: Ten, suit: Heart},
						{rank: Five, suit: Spade},
						{rank: Two, suit: Spade},
					},
					ranks: [5]Rank{
						Ace,
						Queen,
						Ten,
						Five,
						Two,
					},
					handRank: HighCard,
					count: map[Rank]int{
						Ace:   1,
						Queen: 1,
						Ten:   1,
						Five:  1,
						Two:   1,
					},
				},
				{
					cards: [5]Card{
						{rank: Ace, suit: Heart},
						{rank: Queen, suit: Club},
						{rank: Ten, suit: Club},
						{rank: Five, suit: Spade},
						{rank: Two, suit: Diamond},
					},
					ranks: [5]Rank{
						Ace,
						Queen,
						Ten,
						Five,
						Two,
					},
					handRank: HighCard,
					count: map[Rank]int{
						Ace:   1,
						Queen: 1,
						Ten:   1,
						Five:  1,
						Two:   1,
					},
				},
			},
		},
		{
			name: "Pair winner",
			args: args{
				hands: []Hand{
					{
						cards: [5]Card{
							{rank: Ace, suit: Spade},
							{rank: Two, suit: Spade},
							{rank: Queen, suit: Spade},
							{rank: Five, suit: Spade},
							{rank: Ten, suit: Heart},
						},
					},
					{
						cards: [5]Card{
							{rank: Ace, suit: Heart},
							{rank: Two, suit: Diamond},
							{rank: Queen, suit: Club},
							{rank: Five, suit: Spade},
							{rank: Two, suit: Club},
						},
					},
				},
			},
			want: []Hand{
				{
					cards: [5]Card{
						{rank: Ace, suit: Heart},
						{rank: Queen, suit: Club},
						{rank: Five, suit: Spade},
						{rank: Two, suit: Diamond},
						{rank: Two, suit: Club},
					},
					ranks: [5]Rank{
						Two,
						Ace,
						Queen,
						Five,
					},
					handRank: Pair,
					count: map[Rank]int{
						Two:   2,
						Ace:   1,
						Queen: 1,
						Five:  1,
					},
				},
			},
		},
		{
			name: "Pair one winner of two",
			args: args{
				hands: []Hand{
					{
						cards: [5]Card{
							{rank: Ace, suit: Spade},
							{rank: Two, suit: Spade},
							{rank: Queen, suit: Spade},
							{rank: Three, suit: Spade},
							{rank: Two, suit: Heart},
						},
					},
					{
						cards: [5]Card{
							{rank: Ace, suit: Heart},
							{rank: Two, suit: Diamond},
							{rank: Queen, suit: Club},
							{rank: Five, suit: Spade},
							{rank: Two, suit: Club},
						},
					},
				},
			},
			want: []Hand{
				{
					cards: [5]Card{
						{rank: Ace, suit: Heart},
						{rank: Queen, suit: Club},
						{rank: Five, suit: Spade},
						{rank: Two, suit: Diamond},
						{rank: Two, suit: Club},
					},
					ranks: [5]Rank{
						Two,
						Ace,
						Queen,
						Five,
					},
					handRank: Pair,
					count: map[Rank]int{
						Two:   2,
						Ace:   1,
						Queen: 1,
						Five:  1,
					},
				},
			},
		},
		{
			name: "Pair two winners",
			args: args{
				hands: []Hand{
					{
						cards: [5]Card{
							{rank: Ace, suit: Spade},
							{rank: Two, suit: Spade},
							{rank: Queen, suit: Spade},
							{rank: Five, suit: Spade},
							{rank: Two, suit: Heart},
						},
					},
					{
						cards: [5]Card{
							{rank: Ace, suit: Heart},
							{rank: Two, suit: Diamond},
							{rank: Queen, suit: Club},
							{rank: Five, suit: Club},
							{rank: Two, suit: Club},
						},
					},
				},
			},
			want: []Hand{
				{
					cards: [5]Card{
						{rank: Ace, suit: Spade},
						{rank: Queen, suit: Spade},
						{rank: Five, suit: Spade},
						{rank: Two, suit: Spade},
						{rank: Two, suit: Heart},
					},
					ranks: [5]Rank{
						Two,
						Ace,
						Queen,
						Five,
					},
					handRank: Pair,
					count: map[Rank]int{
						Two:   2,
						Ace:   1,
						Queen: 1,
						Five:  1,
					},
				},
				{
					cards: [5]Card{
						{rank: Ace, suit: Heart},
						{rank: Queen, suit: Club},
						{rank: Five, suit: Club},
						{rank: Two, suit: Diamond},
						{rank: Two, suit: Club},
					},
					ranks: [5]Rank{
						Two,
						Ace,
						Queen,
						Five,
					},
					handRank: Pair,
					count: map[Rank]int{
						Two:   2,
						Ace:   1,
						Queen: 1,
						Five:  1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := play(tt.args.hands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("play() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_poker(t *testing.T) {
	type args struct {
		numHands int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "poker 2",
			args: args{
				numHands: 2,
			},
			wantErr: false,
		},
		{
			name: "poker 5",
			args: args{
				numHands: 5,
			},
			wantErr: false,
		},
		{
			name: "poker 10",
			args: args{
				numHands: 10,
			},
			wantErr: false,
		},
		{
			name: "poker 13",
			args: args{
				numHands: 13,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := poker(tt.args.numHands); (err != nil) != tt.wantErr {
				t.Errorf("poker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkPoker(b *testing.B) {
	nums := []struct {
		input int
	}{
		{input: 2},
		{input: 5},
		{input: 7},
		{input: 10},
	}
	for _, num := range nums {
		b.Run(fmt.Sprintf("poker %d", num.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := poker(num.input); err != nil {
					b.Error(err)
				}
			}
		})
	}
}
