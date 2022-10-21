package database

import (
	"finder-rest-api/app/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SearchDataInZincSearch(qs models.QuerySearch) ([]byte, error) {

	qsParameter, err := qs.Serialize()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	url := fmt.Sprintf(os.Getenv("ZINC_SEARCH_API_PATH"), os.Getenv("INDEX_NAME"))

	req, err := http.NewRequest("POST", url, strings.NewReader(qsParameter))
	if err != nil {
		fmt.Println(err)
	}

	req.SetBasicAuth(os.Getenv("ZINC_SEARCH_API_USER"), os.Getenv("ZINC_SEARCH_API_PASS"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "dpacheco")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return nil, fmt.Errorf(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
