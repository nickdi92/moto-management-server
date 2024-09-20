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

	err = http.ListenAndServe(server.Addr, logRequest())
	return server, err
}

func logRequest() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []byte(`{"status": "Go Server is working !!!"}`)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
		w.Write(resp)
		white := color.New(color.FgWhite)
		infoStr := white.Add(color.BgHiMagenta)

		_, _ = infoStr.Println(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		//handler.ServeHTTP(w, r)
	})
}
