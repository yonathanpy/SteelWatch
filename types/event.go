package types

type Event struct {
	IP          string `json:"ip"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	UserAgent   string `json:"user_agent"`
	Score       int    `json:"score"`
	Fingerprint string `json:"fingerprint"`
	Time        string `json:"time"`
}
