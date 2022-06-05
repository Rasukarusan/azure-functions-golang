package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SampleTimerTrigger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("^^^^^^^^^^ SampleTimerTrigger is executed!! ^^^^^^^^^^^^")

	res := response{http.StatusOK, "ok"}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
