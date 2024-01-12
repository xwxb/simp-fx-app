package handler

import (
	"github.com/xwxb/simp-fx-server/util/log"
	"go.uber.org/zap"
	"io"
	"net/http"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		log.Warn("Error copying request body", zap.Error(err))
	}
	log.Info("Echoed request body")
}
