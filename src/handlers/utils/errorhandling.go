package utils

import (
	"bytes"
	"constants"
	"encoding/json"
	"models"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

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

func GetSuccRespObject(v interface{}) events.APIGatewayProxyResponse {
	return GetRespObject(http.StatusOK, v)
}

func GetServErrRespObject(errorText string) events.APIGatewayProxyResponse {
	return GetRespObject(http.StatusInternalServerError,
		models.APIErrResponse{Status: constants.FAILURE,
			Description: errorText})
}

func GetCliErrRespObject(status int, errorText string) events.APIGatewayProxyResponse {
	return GetRespObject(status,
		models.APIErrResponse{Status: constants.FAILURE,
			Description: errorText})
}

func SendErrorResponse(errMap models.ErrorResponse) events.APIGatewayProxyResponse {
	return GetRespObject(
		errMap.HTTPStatusCode,
		models.APIErrorResponse{
			Status:      constants.FAILURE,
			ErrorCode:   errMap.ErrCode,
			Description: errMap.ErrMsg,
		})
}

func SendErrorResponseWithErrorCode(errCode string, errMap map[string]models.ErrorResponse) events.APIGatewayProxyResponse {
	if val, ok := errMap[errCode]; ok {
		return SendErrorResponse(val)
	}
	// check whether it is in the common error map for Rainmaker
	if val, ok := constants.ESPErrMap[errCode]; ok {
		return SendErrorResponse(val)
	} else {
		// Send 500 Error
		//Invalid error code raised, error code not present in provided map or common map
		return SendErrorResponse(constants.ESPErrMap[constants.INVALID_ERROR_CODE_RAISED])
	}
}

func GetErrMsg(errCode string, errMap map[string]models.ErrorResponse) string {
	if val, ok := errMap[errCode]; ok {
		return val.ErrMsg
	}
	return errCode
}
