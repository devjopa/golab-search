package models

import "encoding/json"

type SourceHit struct {
	Body    string
	Date    string
	From    string
	Subject string
	To      string
}

type HitDetail struct {
	Id     string      `json:"_id"`
	Score  json.Number `json:"_score"`
	Source SourceHit   `json:"_source"`
}

type HitHeaderTotal struct {
	Value int64
}

type HitHeader struct {
	Total     HitHeaderTotal
	Max_Score json.Number
	Hits      []HitDetail
}

type ShardDetail struct {
	Total      int
	Successful int
	Skipped    int
	Failed     int
}

type SearchResponse struct {
	Took      int
	Timed_Out bool
	Shards    ShardDetail `json:"_shards"`
	Hits      HitHeader
}
