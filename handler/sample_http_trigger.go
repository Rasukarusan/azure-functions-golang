package handler

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status int
	Rssult string
}

func SampleHttpTrigger(w http.ResponseWriter, r *http.Request) {
	res := response{http.StatusOK, "ok"}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
