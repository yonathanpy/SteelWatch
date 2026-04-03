package core

import "strings"

// Score computes a numeric risk value for suspicious HTTP requests
func Score(ua string, method string, path string) int {
	score := 0
	ua = strings.ToLower(ua)

	// Suspicious user-agents
	if strings.Contains(ua, "sqlmap") {
		score += 60
	}
	if strings.Contains(ua, "nmap") {
		score += 40
	}

	// Odd HTTP methods
	if method != "GET" && method != "POST" {
		score += 20
	}

	// Sensitive paths
	if strings.Contains(path, ".env") || strings.Contains(path, "wp-admin") {
		score += 30
	}

	return score
}
