package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"

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
	result := SimulateGames(simulationInput.Order, simulationInput.Nsimulations)
	//result, opponentTable := SimulateSingleGame(order)

	//log.Printf("Received cards = %v, result = %v, opponent's table = %v", order, result, opponentTable)
	log.Printf("Received cards = %v, result = %v", simulationInput.Order, result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strconv.Itoa(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
