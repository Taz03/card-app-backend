package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetCard(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse {
        StatusCode: 200,
        Body: "get",
    }, nil
}

func main() {
    lambda.Start(GetCard)
}
