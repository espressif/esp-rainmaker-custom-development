package utils

import (
	"constants"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var tagWithTrailingWhitespace = regexp.MustCompile(`^\[.*?\] `)
var tagWithoutTrailingWhitespace = regexp.MustCompile(`^\[.*?\]`)
var squareBraceOpen = "["
var squareBraceClose = "]"

var logger = &log.Logger{
	Out:   os.Stderr,
	Level: log.DebugLevel,
	Formatter: &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	},
}

func init() {
	var logLevel string
	logLevel = GetLogLevel()
	switch logLevel {
	case "Info":
		logger.SetLevel(log.InfoLevel)
	case "Warn":
		logger.SetLevel(log.WarnLevel)
	case "Error":
		logger.SetLevel(log.ErrorLevel)
	case "Debug":
		logger.SetLevel(log.DebugLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
}

/* Following Log functions are supposed take two arguments :
 * RequsetID	(optional) AWSRequestID passed in context object to lambda handler
 * message	(compulsory) Log message to be printed
 */

func AddCallerFunctionInformation(logMessage string) string {
	TAG := "[AddCallerFunctionInformation] "

	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		name := strings.Split(details.Name(), constants.DOT)[1]

		expectedTag := squareBraceOpen + name + squareBraceClose + constants.SPACE

		if tagWithTrailingWhitespace.FindString(logMessage) == expectedTag {
			return logMessage
		} else if tagWithoutTrailingWhitespace.FindString(logMessage) == strings.TrimSpace(expectedTag) {
			logger.Debug(TAG + "DEBUG NOTE: Added missing white space after tag:'" + tagWithoutTrailingWhitespace.FindString(logMessage) + "'")
			logMessage = strings.Replace(logMessage, squareBraceClose, squareBraceClose+constants.SPACE, 1)
		} else {
			logger.Debug(TAG + "DEBUG NOTE: FUNCTION NAME & TAG NAME MISMATCH ! Expected:'" + squareBraceOpen + name + squareBraceClose + "', but got:'" + logMessage[:strings.Index(logMessage, squareBraceClose)+1] + "'")
			logMessage = expectedTag + strings.TrimSpace(logMessage[strings.Index(logMessage, squareBraceClose)+1:])
		}
	}
	return logMessage
}

func LogDebug(args ...string) {
	logMessage := AddCallerFunctionInformation(args[0])
	if len(args) == 2 {
		logger.Debug(logMessage + " " + args[1])
	} else {
		logger.Debug(logMessage)
	}
}

func LogInfo(args ...string) {
	logMessage := AddCallerFunctionInformation(args[0])
	if len(args) == 2 {
		logger.Info(logMessage + " " + args[1])
	} else {
		logger.Info(logMessage)
	}
}

func LogWarn(args ...string) {
	logMessage := AddCallerFunctionInformation(args[0])
	if len(args) == 2 {
		logger.Warn(logMessage + " " + args[1])
	} else {
		logger.Warn(logMessage)
	}
}

func LogError(args ...string) {
	logMessage := AddCallerFunctionInformation(args[0])
	if len(args) == 2 {
		logger.Error(logMessage + " " + args[1])
	} else {
		logger.Error(logMessage)
	}
}
