package swagger

import (
	"log"
	"net/http"
	"encoding/json"
)

type MyResponse struct {
	OkMessage    interface{} `json:"ok,omitempty"`
	ErrorMessage interface{} `json:"error,omitempty"`
}

func Response(response interface{}, w http.ResponseWriter, code int) {
	jsonData, jErr := json.Marshal(&response)

	if jErr != nil {
		log.Fatal(jErr.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(jsonData)
	w.WriteHeader(code)
}
