package models

import (
	"encoding/json"
	"fmt"
)

type Query struct {
	Term       string
	Start_Time string
	End_Time   string
}

type QuerySearch struct {
	Search_Type string `json:"search_type"`
	Query       Query
	From        int64 `json:"from"`
	Max_Results int64 `json:"max_results"`
}

func CreateQuerySearch(term string, from int64, max_Results int64) *QuerySearch {

	return &QuerySearch{
		Search_Type: "match",
		Query:       Query{Term: term, Start_Time: "2000-06-02T14:28:31.894Z", End_Time: "2023-06-02T14:28:31.894Z"},
		From:        from,
		Max_Results: max_Results,
	}

}

func (qs *QuerySearch) Serialize() (string, error) {

	qsJSON, err := json.Marshal(qs)

	if err != nil {
		fmt.Println("Error struct converting to JSON structure", err)
		return "", err
	}

	return string(qsJSON), nil
}
