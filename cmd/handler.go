package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/weather-app/types"
)


func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) 

	data := map[string]string{
		"status":  "ok",
		"env":     "8080",
		"version": "1.1.0",
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}


func (h *application) weatherHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.contextTimeout)*time.Second)
	defer cancel()

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, `{"status":"error", "message":"city params cannot be empty"}`, http.StatusBadRequest)
		return
	}

	response, err := h.weatherService.GetWeatherByCity(ctx, city)
	// When the Redis key expires, It will be skip err and the data is re-stored in Redis.
	// This prevents continuous API requests with the previously stored key in Redis.
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"data": response,
		})
		return
	}

	apiUrl := &types.Api{
		Url: h.config.apiURL,
		City: city,
		ApiKey: h.config.apiKey,
	}

	store, err := h.weatherService.CreateWeather(ctx, apiUrl)
    if err != nil {
        http.Error(w, `{"status":"error","message":"`+err.Error()+`"}`, http.StatusBadRequest)
        return
    }

	 // Return the fetched weather data
	 w.Header().Set("Content-Type", "application/json")
	 json.NewEncoder(w).Encode(map[string]interface{}{
		 "status": "success",
		 "data":   store,
	 })
}
