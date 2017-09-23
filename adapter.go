package notary

import (
	"encoding/json"
	"net/http"

	log "go.uber.org/zap"
)

type response struct {
	Resp interface{} `json:"response"`
}
type Adapter func(http.Handler) http.Handler

// Adapt h with all specified adapters.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// decode can be this simple to start with, but can extended
// later to support differnt formats and behaviours without
// changing the interface
func decode(r *http.Request, v interface{}) error {

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

// TODO:
// Handle read write locking
type Handler struct {
	logger *log.Logger
	db     *dbServer
	data   interface{}
	status int
}

/**
 * Respond to request
 */
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(&response{h.data})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(h.status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
