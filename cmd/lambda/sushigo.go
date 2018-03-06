package main

import (
	"github.com/BasPH/lamarl-go/sushigo"
	"log"
	"math"
	"math/rand"
	"sort"
	"sync"
)

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

	for card, n := range sushigo.SushiGoCards {
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

	numSashimi := countOccurences(thisTable, "sashimi")
	score += int(math.Floor(float64(numSashimi)/3) * 10)

	numTempura := countOccurences(thisTable, "tempura")
	score += int(math.Floor(float64(numTempura)/3) * 10)

	numDumpling := countOccurences(thisTable, "tempura")
	dumplingScores := []int{0, 1, 3, 6, 10, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15}
	score += dumplingScores[numDumpling]

	numTofu := countOccurences(thisTable, "tempura")
	tofuScores := []int{0, 2, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	score += tofuScores[numTofu]

	numEel := countOccurences(thisTable, "tempura")
	eelScores := []int{0, -3, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	score += eelScores[numEel]

	return score
}

// SimulateSingleGame docs
func SimulateSingleGame(cards []string) (int, []string) {
	log.Printf("Player cards = %v", cards)
	cardsOrdered := indexCards(cards)
	log.Printf("Player cards ordered = %v", cardsOrdered)

	handPlayer := generateHand()
	log.Printf("Generated player's hand = %v", handPlayer)
	handOpponent := generateHand()
	log.Printf("Generated opponent's hand  = %v", handOpponent)

	var tablePlayer []string
	var tableOpponent []string

	for len(tablePlayer) < 13 {
		// Shuffle cards
		// Note: You cannot use AWS CodeBuild to build the project because it uses Go 1.7.3.
		// sort.Slice was implemented in 1.8.
		sort.Slice(handPlayer, func(i, j int) bool {
			return cardsOrdered[handPlayer[i]] < cardsOrdered[handPlayer[j]]
		})

		shuffleCards(handOpponent)
		log.Printf("Shuffled opponent's hand  = %v", handOpponent)

		// Put cards on the table
		selectedPlayer := handPlayer[len(handPlayer)-1]
		handPlayer = handPlayer[:len(handPlayer)-1]
		tablePlayer = append(tablePlayer, selectedPlayer)
		log.Printf("Selected, hand & table player = %v, %v, %v", selectedPlayer, handPlayer, tablePlayer)

		selectedOpponent := handOpponent[len(handOpponent)-1]
		handOpponent = handOpponent[:len(handOpponent)-1]
		tableOpponent = append(tableOpponent, selectedOpponent)
		log.Printf("Selected, hand & table opponent = %v, %v, %v", selectedOpponent, handOpponent, tableOpponent)

		// Pass cards to each other
		handPlayer, handOpponent = handOpponent, handPlayer
		log.Printf("Passed cards. Hands are now: player = %v, opponent = %v", handPlayer, handOpponent)
	}

	tablePlayer = append(tablePlayer, handPlayer[len(handPlayer)-1])
	tableOpponent = append(tableOpponent, handOpponent[len(handOpponent)-1])
	log.Printf("Table player: %v", tablePlayer)
	log.Printf("Table opponent: %v", tableOpponent)

	if scoreTable(tablePlayer, tableOpponent) > scoreTable(tableOpponent, tablePlayer) {
		return 1, tableOpponent
	}
	return 0, tableOpponent
}

// SimulateGames docs
func SimulateGames(order []string, nSim int) int {
	var wg sync.WaitGroup
	wg.Add(nSim)
	count := 0

	for i := 0; i < nSim; i++ {
		go func(order []string) {
			defer wg.Done()
			result, _ := SimulateSingleGame(order)
			count += result
		}(order)
	}

	wg.Wait()
	return count
}
