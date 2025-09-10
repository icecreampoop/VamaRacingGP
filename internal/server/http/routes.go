package http

import (
	"net/http"
)

func (h *Handler) NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.defaultHandler)
	mux.HandleFunc("/showusers", h.showAllFromUsersTable)

	return LoggingMiddleware(mux)
}
