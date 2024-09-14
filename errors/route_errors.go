package errors

import "fmt"

type RouteErrorCode string

const RouteErrorCode_InvalidRoute RouteErrorCode = "InvalidRoute"
const RouteErrorCode_NoRoutesRegistered RouteErrorCode = "NoRoutesRegistered"

type RouteError struct {
	Code    RouteErrorCode
	Message string
}

func (r RouteError) Error() string {
	return fmt.Sprintf("RouteError. Code: %s, Message: %s", r.Code, r.Message)
}
