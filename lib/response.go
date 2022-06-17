package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data interface{} `json:"message"`
}

func ResponseBuilder(w http.ResponseWriter, result ...interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if len(result) == 0 {
		result = append(result, "OK")
	}

	m := Response{
		Data: result[0],
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error: %v", err)
	}
}
