package server

import (
	"fmt"
	"net/http"
)

var LoginRoute = func(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Login Route! @TODO")
}
