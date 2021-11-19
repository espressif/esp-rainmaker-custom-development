package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"handlers/utils"
	"models"
	"net/http"
)

func fetch(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	TAG := "[fetch] "
	utils.LogDebug(TAG + "In Get Func")
	//Add logic here
	return utils.GetRespObject(http.StatusOK, models.APIResponse{"Success", "Output from GET func"}), nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	TAG := "[create] "
	utils.LogDebug(TAG + "In Get Func")
	//Add logic here
	return utils.GetRespObject(http.StatusOK, models.APIResponse{"Success", "Output from POST func"}), nil
}

func update(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	TAG := "[update] "
	utils.LogDebug(TAG + "In Get Func")
	//Add logic here
	return utils.GetRespObject(http.StatusOK, models.APIResponse{"Success", "Output from PUT func"}), nil
}

func delete(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	TAG := "[delete] "
	utils.LogDebug(TAG + "In Get Func")
	//Add logic here
	return utils.GetRespObject(http.StatusOK, models.APIResponse{"Success", "Output from DELETE func"}), nil
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return fetch(req)
	case "POST":
		return create(req)
	case "PUT":
		return update(req)
	case "DELETE":
		return delete(req)
	default:
		return events.APIGatewayProxyResponse{}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
