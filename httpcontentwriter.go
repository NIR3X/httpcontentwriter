package httpcontentwriter

import (
	"net/http"
	"sync"
)

type HttpContentWriter struct {
	mtx          sync.Mutex
	isFirstWrite bool
	w            http.ResponseWriter
}

func NewHttpContentWriter(w http.ResponseWriter) *HttpContentWriter {
	return &HttpContentWriter{
		mtx:          sync.Mutex{},
		isFirstWrite: true,
		w:            w,
	}
}

func (h *HttpContentWriter) Write(p []uint8) (n int, err error) {
	isFirstWrite := false
	h.mtx.Lock()
	if h.isFirstWrite {
		h.isFirstWrite = false
		isFirstWrite = true
	}
	h.mtx.Unlock()
	if isFirstWrite {
		h.w.Header().Set("Content-Type", http.DetectContentType(p))
	}
	return h.w.Write(p)
}
