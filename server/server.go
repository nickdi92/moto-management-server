package server

import (
	"fmt"
	"moto-management-server/business_logic"
	"moto-management-server/utils"
	"net/http"

	"github.com/fatih/color"
)

func (s *MotoManagementServer) NewMotoManagementServer() (*MotoManagementServer, error) {
	bl := &business_logic.BusinessLogic{}
	server := &MotoManagementServer{
		Addr:          ":8080",
		businessLogic: bl.NewBusinessLogic(),
	}

	server.RegisterRoutes()

	err := server.HandleRoutes()
	if err != nil {
		return nil, err
	}

	utils.SuccessOutput("Webserver started")
	utils.InfoOutput(fmt.Sprintf("Listening on localhost%s", server.Addr))

	err = http.ListenAndServe(server.Addr, logRequest(http.DefaultServeMux))
	return server, err
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Connection", "keep-alive")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Add("Access-Control-Max-Age", "86400")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		white := color.New(color.FgWhite)
		infoStr := white.Add(color.BgHiMagenta)

		_, _ = infoStr.Println(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}
