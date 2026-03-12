package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

// Request structure for incoming JSON
type CalcRequest struct {
	Number float64 `json:"number"`
	Power  float64 `json:"power"`
}

// Response structure for outgoing JSON
type CalcResponse struct {
	Result float64 `json:"result"`
}

func main() {
	// 1. Serve static files (AngularJS frontend)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 2. API Endpoints for the 4 required functions
	http.HandleFunc("/api/sqrt", func(w http.ResponseWriter, r *http.Request) {
		var req CalcRequest
		json.NewDecoder(r.Body).Decode(&req)
		json.NewEncoder(w).Encode(CalcResponse{Result: math.Sqrt(req.Number)})
	})

	http.HandleFunc("/api/factorial", func(w http.ResponseWriter, r *http.Request) {
		var req CalcRequest
		json.NewDecoder(r.Body).Decode(&req)
		res := 1.0
		for i := 1.0; i <= req.Number; i++ { res *= i }
		json.NewEncoder(w).Encode(CalcResponse{Result: res})
	})

	http.HandleFunc("/api/ln", func(w http.ResponseWriter, r *http.Request) {
		var req CalcRequest
		json.NewDecoder(r.Body).Decode(&req)
		json.NewEncoder(w).Encode(CalcResponse{Result: math.Log(req.Number)})
	})

	http.HandleFunc("/api/power", func(w http.ResponseWriter, r *http.Request) {
		var req CalcRequest
		json.NewDecoder(r.Body).Decode(&req)
		json.NewEncoder(w).Encode(CalcResponse{Result: math.Pow(req.Number, req.Power)})
	})

	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}