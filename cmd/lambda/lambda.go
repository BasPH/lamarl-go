package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"

	"fmt"
	"github.com/BasPH/lamarl-go/sushigo"
)

// Response docs
type Response struct {
	Result int `json:"result"`
}

// Handler docs
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	bytes := []byte(request.Body)
	var simulationInput sushigo.SimulationInput
	json.Unmarshal(bytes, &simulationInput)
	if !simulationInput.ValidateCards() {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("invalid cards")
	}

	result := SimulateGames(simulationInput.Order, simulationInput.Nsimulations)
	log.Printf("Received cards = %v, result = %v", simulationInput.Order, result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strconv.Itoa(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
