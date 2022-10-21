package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func SendDataToZincSearch(data string) error {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		return err
	}

	url := fmt.Sprintf(os.Getenv("ZINC_SEARCH_API_PATH"), os.Getenv("INDEX_NAME"))

	req, err := http.NewRequest("POST", url, strings.NewReader(data))

	if err != nil {
		fmt.Println(err)
		return err
	}

	req.SetBasicAuth(os.Getenv("ZINC_SEARCH_API_USER"), os.Getenv("ZINC_SEARCH_API_PASS"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "dpacheco")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return fmt.Errorf(resp.Status)
	}

	return nil
}
