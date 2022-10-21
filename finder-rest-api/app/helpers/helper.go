package helpers

import (
	"encoding/json"
	"errors"
	"finder-rest-api/app/models"
	"log"
	"net/http"
)

func GenerateJSONResponse(data []byte) (*models.SearchResponse, error) {
	var searchResponseObject *models.SearchResponse
	err := json.Unmarshal(data, &searchResponseObject)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return searchResponseObject, nil

}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func GenerateJSONTokenResponse(data []byte) (*models.UserTokenResponse, error) {
	var userResponseObject *models.UserTokenResponse
	err := json.Unmarshal(data, &userResponseObject)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return userResponseObject, nil

}
