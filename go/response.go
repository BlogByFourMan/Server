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

func Options(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials, Access-Control-Allow-Methods, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

func Response(response interface{}, w http.ResponseWriter, code int) {
	jsonData, jErr := json.Marshal(&response)

	if jErr != nil {
		log.Fatal(jErr.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials, Access-Control-Allow-Methods, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
	w.WriteHeader(code)
}
