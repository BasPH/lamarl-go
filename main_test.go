package main_test

import (
	"github.com/BasPH/lamarl-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request main.SimulationInput
		expect  string
		err     error
	}{
		{
			request: main.SimulationInput{20, []string{"1", "2", "3"}},
			expect:  "{\"nsimulations\":20,\"order\":[\"1\",\"2\",\"3\"]}",
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
