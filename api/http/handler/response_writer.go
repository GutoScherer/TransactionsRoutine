package handler

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	writer http.ResponseWriter
}

func newResponseWriter(rw http.ResponseWriter) *responseWriter {
	return &responseWriter{writer: rw}
}

func (rw responseWriter) outputResponse(statusCode int, payload interface{}) {
	rw.writer.Header().Set("Content-Type", "application/json")
	rw.writer.WriteHeader(statusCode)

	if payload != nil {
		json.NewEncoder(rw.writer).Encode(payload)
	}
}
