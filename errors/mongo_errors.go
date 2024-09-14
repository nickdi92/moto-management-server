package errors

import "fmt"

type MongoErrorCode string

const MongoErrorCode_FailedToConnect = "FailedToConnect"
const MongoErrorCode_FailedToDisconnect = "FailedToDisconnect"
const MongoErrorCode_FailedToPing = "FailedToPing"

type MongoErrors struct {
	Code    MongoErrorCode
	Message string
}

func (m MongoErrors) Error() string {
	return fmt.Sprintf("Exit with code: %s. Error is: %v", m.Code, m.Message)
}
