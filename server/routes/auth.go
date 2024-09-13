package routes

import (
	"fmt"
	"net/http"
)

var AuthRoute = func(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Auth Route! @TODO")
}
