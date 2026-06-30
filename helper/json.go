package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	encoder := json.NewDecoder(request.Body)
	err := encoder.Decode(result)
	PanicErr(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-type", "Application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicErr(err)
}
