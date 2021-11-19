package utils

import (
	"os"
)

var ACCOUNT_ID = "ACCOUNT_ID"
var REGION = "REGION"
var LOG_LEVEL = "LOG_LEVEL"

func GetAccountId() string {
	return os.Getenv(ACCOUNT_ID)
}

func GetRegion() string {
	return os.Getenv(REGION)
}

func GetLogLevel() string {
	return os.Getenv(LOG_LEVEL)
}
