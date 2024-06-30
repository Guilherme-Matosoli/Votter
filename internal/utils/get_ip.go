package utils

import "net/http"

func GetIp(r *http.Request) string {
	xReal := r.Header.Get("X-Real-Ip")
	if xReal != "" {
		return xReal
	}

	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}

	return r.RemoteAddr
}
