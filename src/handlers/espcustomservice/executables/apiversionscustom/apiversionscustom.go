package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type APIResponse struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

func GetRespObject(status int, v interface{}) events.APIGatewayProxyResponse {
	body, err := json.Marshal(v)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "JSON Marshaling Failure",
		}
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers:         map[string]string{"Access-Control-Allow-Origin": "*", "X-Content-Type-Options": "nosniff"},
	}
}

func fetch(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return GetRespObject(http.StatusOK, APIResponse{"Success", "return from get api versions func"}), nil
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return fetch(req)
	default:
		return events.APIGatewayProxyResponse{}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
