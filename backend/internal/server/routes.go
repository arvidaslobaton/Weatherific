package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Server represents the HTTP server

// Routes returns the handler for the server's routes
func (s *Server) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.Welcome).Methods("GET")
	r.HandleFunc("/getWeather", s.getWeather).Methods("POST")

	// CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Wrap the router with CORS handler
	handler := c.Handler(r)
	http.Handle("/", r)
	return handler
}

func (s *Server) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Terrific Weather!")
}

type WeatherRequest struct {
	Name string `json:"name"`
}

type WeatherResponse struct {
	 Cod       	interface{} `json:"cod"`
	Name      	string  	`json:"name"`
	Temp      	float64 	`json:"temp"`
	Weather   	string  	`json:"weather"`
	FeelsLike 	float64 	`json:"feels_like"`
	Humidity  	float64 	`json:"humidity"`
	WindSpeed 	float64 	`json:"wind_speed"`
}

func (s *Server) getWeather(w http.ResponseWriter, r *http.Request) {
	var req WeatherRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("OWM_API_KEY")

	weather, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		http.Error(w, "Failed to initialize weather data", http.StatusInternalServerError)
		return
	}

	err = weather.CurrentByName(req.Name)
	if err != nil {
		http.Error(w, "Failed to get weather data", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	resp := WeatherResponse{
		Cod:       weather.Cod,
		Name:      weather.Name,
		Weather:   weather.Weather[0].Main,
		Temp:      weather.Main.Temp,
		FeelsLike: weather.Main.FeelsLike,
		Humidity:  float64(weather.Main.Humidity),
		WindSpeed: weather.Wind.Speed,
	}

	weatherData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal weather data", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(weatherData)
}
