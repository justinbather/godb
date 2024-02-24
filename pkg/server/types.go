package server

type requestItem struct {
	Value   interface{} `json:"value"`
	Key     string      `json:"key"`
	TTL     int         `json:"ttl"`
	Sliding bool        `json:"sliding"`
}