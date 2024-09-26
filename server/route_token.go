package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"moto-management-server/utils"
	"net/http"
	"strings"
)

var RefreshTokenRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var tokenRequest models.TokenRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal([]byte(body), &tokenRequest)

	validationErr := s.ValidateRequest(tokenRequest)
	if validationErr != nil {
		err := map[string]interface{}{"RefreshTokenRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	tokenRequest.Username = strings.ToLower(tokenRequest.Username)
	tokenRequest.Password, _ = utils.Password(tokenRequest.Password).Hash()
	token := s.token.NewToken(tokenRequest.Username, tokenRequest.Password)
	token.RefreshToken()

	blUser, err := s.businessLogic.GetUserByUsername(tokenRequest.Username)
	if err != nil {
		err := map[string]interface{}{"RefreshTokenRoute": err.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	blUser.Token = token.Token
	blUser.ExpireAt = token.ExpiresAt

	_, updateErr := s.businessLogic.UpdateUser(blUser)
	if updateErr != nil {
		err := map[string]interface{}{"RefreshTokenRoute": updateErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	tokenResp := models.TokenResponse{
		StatusCode: http.StatusOK,
		Token:      token.Token,
		ExpireAt:   token.ExpiresAt,
	}

	s.HandleResponse(writer, tokenResp)
}
