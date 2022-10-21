package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"finder-rest-api/app/database"
	"finder-rest-api/app/helpers"
	"finder-rest-api/app/models"

	"github.com/go-chi/chi/v5"
)

func (r *Route) CreateSearchHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		termParam := chi.URLParam(r, "term")
		fromParam, err := strconv.ParseInt(chi.URLParam(r, "from"), 0, 32)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		maxResultsParam, err := strconv.ParseInt(chi.URLParam(r, "maxResults"), 0, 32)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		objQuerySearch := models.CreateQuerySearch(termParam, fromParam, maxResultsParam)

		dataResult, err := database.SearchDataInZincSearch(*objQuerySearch)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error getting data %s-%d-%d", termParam, fromParam, maxResultsParam)))
			return
		}

		resp, err := helpers.GenerateJSONResponse(dataResult)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf(err.Error())))
			return
		}

		helpers.SendResponse(w, r, resp, http.StatusOK)

	}
}

func (r *Route) getToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var u models.UserToken

		hasError := false
		errMessage := ""
		token := ""

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}

		if u.Username == "" || u.Password == "" {
			hasError = true
			errMessage = "Missing User or Password"
		}

		if u.Username != os.Getenv("USER_JWT") || u.Password != os.Getenv("PASS_JWT") {
			hasError = true
			errMessage = "User or Password are invalid"
		}

		if !hasError {
			token = MakeToken(u.Username)
		}

		resp := models.CreateUserTokenResponse(hasError, errMessage, token)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf(err.Error())))
			return
		}

		helpers.SendResponse(w, r, resp, http.StatusOK)

	}
}

func MakeToken(name string) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"username": name})
	return tokenString
}
