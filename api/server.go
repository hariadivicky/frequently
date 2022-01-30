package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hariadivicky/frequently/appearance"
)

func StartServer(address string, maxUpload int64) error {
	server := http.Server{
		Handler:      &serverHandler{maxUpload: maxUpload},
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  60 * time.Second,
		Addr:         address,
	}

	log.Printf("server listening on %s", address)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

type serverHandler struct {
	maxUpload int64
}

func (h *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errResponse(w, "method not allowed", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(h.maxUpload)

	uploaded, _, err := r.FormFile("file")
	if err != nil {
		h.errResponse(w, "invalid file", http.StatusBadRequest)
		return
	}
	defer uploaded.Close()

	data, err := ioutil.ReadAll(uploaded)
	if err != nil {
		h.errResponse(w, "unreadable file", http.StatusBadRequest)
		return
	}

	if insensitive := r.FormValue("insensitive"); insensitive != "" {
		data = bytes.ToLower(data)
	}

	var max int64 = 10
	maxVal := r.FormValue("max")
	if maxVal != "" {
		converted, err := strconv.ParseInt(maxVal, 10, 32)
		if err != nil {
			h.errResponse(w, "invalid integer value for max_result", http.StatusBadRequest)
			return
		}

		max = converted
	}

	counter := appearance.NewCounter()
	result := counter.Top(data, int(max))

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h *serverHandler) errResponse(w http.ResponseWriter, msg string, status int) {
	wrapped := map[string]string{
		"error": msg,
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(wrapped); err != nil {
		log.Printf("error encoding json response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
