package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var response = make(map[string]string)
	now := time.Now()
	w.Header().Add("Content-Type", "application/json")
	if val, ok := vars["tz"]; ok {
		timezones := strings.Split(val, ",")
		for _, tz := range timezones {
			locale, err := time.LoadLocation(tz)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"error":       "timezone not found",
					"error_value": tz,
				})
				return
			}
			response[tz] = now.In(locale).String()
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response["current_time"] = now.UTC().String()
		json.NewEncoder(w).Encode(response)
	}
}
