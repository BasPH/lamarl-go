package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

// SimulationInput docs
type SimulationInput struct {
	Order []string `json:"order"`
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
	order := simulationInput.Order
	result := SimulateSingleGame(order)

	log.Printf("Received cards = %v, result = %v", order, result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strconv.Itoa(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
