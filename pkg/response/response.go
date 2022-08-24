package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type baseResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Error error `json:"error"`
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data := &baseResponse{
		Data: payload,
	}
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
}

func ResponseWithNoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func ResponseError(w http.ResponseWriter, code int, e error) {
	data := &errorResponse{
		Error: e,
	}
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
}
