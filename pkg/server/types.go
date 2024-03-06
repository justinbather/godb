package server

type requestItem struct {
	Value   string `json:"value"`
	Key     string `json:"key"`
	TTL     int    `json:"ttl"`
	Sliding bool   `json:"sliding"`
}

type Config struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}
