package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{Body: "{\"nsimulations\":20,\"order\":[\"1\",\"2\",\"3\"]}"},
			expect:  "20",
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
