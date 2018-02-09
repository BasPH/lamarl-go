package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"fmt"
	"strconv"
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
	result := simulationInput.Nsimulations
	fmt.Println(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strconv.Itoa(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
