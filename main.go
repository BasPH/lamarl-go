package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
)

// SimulationInput docs
type SimulationInput struct {
	Nsimulations int      `json:"nsimulations"`
	Order        []string `json:"order"`
}

// Response docs
type Response struct {
	Result int `json:"result"`
}

// Handler docs
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	bytes := []byte(request.Body)
	var simulationInput SimulationInput
	json.Unmarshal(bytes, &simulationInput)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(simulationInput.Nsimulations),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
