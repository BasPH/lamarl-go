package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)

// SimulationInput docs
type SimulationInput struct {
	Order        []string `json:"order"`
	Nsimulations int      `json:"nsimulations"`
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
