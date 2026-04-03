package handlers

import (
	"net"
	"net/http"
	"time"
	"steelwatch/core"
	"steelwatch/storage"
	"steelwatch/types"
)

type HTTPHandler struct {
	Store *storage.Store
}

func getIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func (h *HTTPHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	raw := ip + r.UserAgent() + r.URL.Path
	event := types.Event{
		IP:          ip,
		Path:        r.URL.Path,
		Method:      r.Method,
		UserAgent:   r.UserAgent(),
		Score:       core.Score(r.UserAgent(), r.Method, r.URL.Path),
		Fingerprint: core.Fingerprint(raw),
		Time:        time.Now().Format(time.RFC3339),
	}

	h.Store.Add(event)
	w.Write([]byte("OK"))
}
