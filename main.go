package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
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

func handler(request SimulationInput) (events.APIGatewayProxyResponse, error) {

	b, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}
