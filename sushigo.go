package main

import (
	"math/rand"
)

var sushiGoCards = map[string](int){
	"tempura":       14,
	"sashimi":       14,
	"dumpling":      14,
	"maki-2":        12,
	"maki-3":        8,
	"maki-1":        6,
	"salmon-nigiri": 10,
	"squid-nigiri":  5,
	"egg-nigiri":    5,
	"pudding":       10,
	"wasabi":        6,
	"chopsticks":    4,
}

func indexCards(cards []string) map[string]int {
	cardIndex := make(map[string]int, len(cards))
	for i, card := range cards {
		cardIndex[card] = i
	}
	return cardIndex
}

func shuffleCards(cards []string) []string {
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

// We're only doing 2 player games now, so 10 cards each.
// 2 players -> 10 cards each.
// 3 players -> 9 cards each.
// 4 players -> 8 cards each.
// 5 players -> 7 cards each.
func generateHand() []string {
	var cards []string

	for card, n := range sushiGoCards {
		for i := 0; i < n; i++ {
			cards = append(cards, card)
		}
	}

	return shuffleCards(cards)[:10]
}

// Count number of occurences of something in a slice
func countOccurences(cards []string, thing string) int {
	count := 0
	for _, item := range cards {
		if item == thing {
			count++
		}
	}
	return count
}
