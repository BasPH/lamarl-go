package main

import (
	"testing"
)

func TestCountOccurences(t *testing.T) {
	a := []string{"a", "b", "c", "a", "a"}
	n := countOccurences(a, "a")
	if n != 3 {
		t.Error("Expected 3 but got", n)
	}
}

func benchmarkGenerateHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateHand()
	}
}

func BenchmarkGenerateHand1(b *testing.B)    { benchmarkGenerateHand(b) }
func BenchmarkGenerateHand10(b *testing.B)   { benchmarkGenerateHand(b) }
func BenchmarkGenerateHand100(b *testing.B)  { benchmarkGenerateHand(b) }
func BenchmarkGenerateHand1000(b *testing.B) { benchmarkGenerateHand(b) }

func benchmarkShuffle(b *testing.B) {
	cards := generateHand()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		shuffleCards(cards)
	}
}

func BenchmarkShuffle1(b *testing.B)    { benchmarkShuffle(b) }
func BenchmarkShuffle10(b *testing.B)   { benchmarkShuffle(b) }
func BenchmarkShuffle100(b *testing.B)  { benchmarkShuffle(b) }
func BenchmarkShuffle1000(b *testing.B) { benchmarkShuffle(b) }

func BenchmarkSimulateGames(b *testing.B) {
	cards := generateHand()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SimulateGames(cards, n)
	}
}

func BenchmarkSimulateGames1(b *testing.B)    { benchmarkShuffle(b) }
func BenchmarkSimulateGames10(b *testing.B)   { benchmarkShuffle(b) }
func BenchmarkSimulateGames100(b *testing.B)  { benchmarkShuffle(b) }
func BenchmarkSimulateGames1000(b *testing.B) { benchmarkShuffle(b) }
