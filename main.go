package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

type SimulationInput struct {
	Nsimulations int      `json:"nsimulations"`
	Order        []string `json:"order"`
}

type Response struct {
	Result int `json:"result"`
}

func handler(request SimulationInput) (events.APIGatewayProxyResponse, error) {

	fmt.Println(request)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello!!",
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}
