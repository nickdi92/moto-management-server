package errors

import "fmt"

type RouteErrorCode string

const RouteErrorCode_InvalidRoute RouteErrorCode = "InvalidRoute"
const RouteErrorCode_NoRoutesRegistered RouteErrorCode = "NoRoutesRegistered"

type RouteErrors struct {
	Code    RouteErrorCode
	Message string
}

func (e RouteErrors) Error() string {
	return fmt.Sprintf("RouteError. Code: %s, Message: %s", e.Code, e.Message)
}
