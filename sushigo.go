package main

import (
	"math/rand"
	"sort"
	"math"
)

var sushiGoCards = map[string](int){
	"maki-1":   5,
	"maki-2":   5,
	"maki-3":   5,
	"sashimi":  5,
	"egg":      5,
	"salmon":   5,
	"squid":    5,
	"wasabi":   5,
	"pudding":  5,
	"tempura":  5,
	"dumpling": 5,
	"tofu":     5,
	"eel":      5,
	"temaki":   5,
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

	return shuffleCards(cards)[:14]
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

func scoreTable(thisTable []string, thatTable []string) int {
	score := 0

	thisTemaki := countOccurences(thisTable, "temaki")
	thatTemaki := countOccurences(thatTable, "temaki")
	if thisTemaki > thatTemaki {
		score += 4
	}

	multiplier := 1
	for _, card := range thisTable {
		if card == "wasabi" {
			multiplier = 2
		}
		if card == "egg" {
			score += multiplier
			multiplier = 1
		}
		if card == "salmon" {
			score += multiplier * 2
			multiplier = 1
		}
		if card == "squid" {
			score += multiplier * 3
			multiplier = 1
		}
	}

	thisMaki := countOccurences(thisTable, "maki-1") + countOccurences(thisTable, "maki-2")*2 + countOccurences(thisTable, "maki-3")*3
	thatMaki := countOccurences(thatTable, "maki-1") + countOccurences(thatTable, "maki-2")*2 + countOccurences(thatTable, "maki-3")*3
	if thisMaki > thatMaki {
		score += 6
	} else if thisMaki > 1 {
		score += 3
	}

	num_sashimi := countOccurences(thisTable, "sashimi")
	score += int(math.Floor(float64(num_sashimi)/3) * 10)

	num_tempura := countOccurences(thisTable, "tempura")
	score += int(math.Floor(float64(num_tempura)/3) * 10)

	num_dumpling := countOccurences(thisTable, "tempura")
	dumpling_scores := []int{0, 1, 3, 6, 10, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15}
	score += dumpling_scores[num_dumpling]

	num_tofu := countOccurences(thisTable, "tempura")
	tofu_scores := []int{0, 2, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	score += tofu_scores[num_tofu]

	num_eel := countOccurences(thisTable, "tempura")
	eel_scores := []int{0, -3, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	score += eel_scores[num_eel]

	return score
}

func SimulateSingleGame(cards []string) int {
	cardsOrdered := indexCards(cards)

	handPlayer := generateHand()
	handOpponent := generateHand()

	var tablePlayer []string
	var tableOpponent []string

	for len(tablePlayer) < 13 {
		// Shuffle cards
		sort.Slice(handPlayer, func(i, j int) bool {
			return cardsOrdered[handPlayer[i]] < cardsOrdered[handPlayer[j]]
		})
		shuffleCards(handOpponent)

		// Put cards on the table
		selectedPlayer := handPlayer[len(handPlayer)-1]
		handPlayer = handPlayer[:len(handPlayer)-1]
		tablePlayer = append(tablePlayer, selectedPlayer)
		selectedOpponent := handOpponent[len(handOpponent)-1]
		handOpponent = handOpponent[:len(handOpponent)-1]
		tableOpponent = append(tableOpponent, selectedOpponent)

		// Pass cards to each other
		handPlayer, handOpponent = handOpponent, handPlayer
	}

	tablePlayer = append(tablePlayer, handPlayer[len(handPlayer)-1])
	tableOpponent = append(tableOpponent, handOpponent[len(handOpponent)-1])

	tablePlayer = append(tablePlayer, handPlayer[len(handPlayer)-1])
	tableOpponent = append(tableOpponent, handOpponent[len(handOpponent)-1])
	//fmt.Printf("table_player: %v\n", table_player)
	//fmt.Printf("table_other:  %v\n", table_other)
	if scoreTable(tablePlayer, tableOpponent) > scoreTable(tableOpponent, tablePlayer) {
		return 1
	}
	return 0
}

func SimulateGames(order []string, n_sim int) int {
	count := 0
	for i := 1; i <= n_sim; i++ {
		count += SimulateSingleGame(order)
	}
	return count
}

//func main() {
//	order := []string{"maki-1", "maki-2", "maki-3", "sashimi", "egg", "salmon", "squid", "wasabi", "pudding", "tempura", "dumpling", "tofu", "eel", "temaki"}
//	SimulateSingleGame(order)
//	start := time.Now()
//	fmt.Println(SimulateGames(order, 200000))
//	fmt.Printf("time taken: %s", time.Since(start))
//}
