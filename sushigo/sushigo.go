package sushigo

type SimulationInput struct {
	Nsimulations int      `json:"nsimulations"`
	Order        []string `json:"order"`
}

var SushiGoCards = map[string]int{
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

func (si SimulationInput) ValidateCards() bool {
	for _, c := range si.Order {
		if _, exists := SushiGoCards[c]; !exists {
			return false
		}
	}
	return true
}
