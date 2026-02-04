package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Payload struct {
	Message  string `json:"message"`
	Password string `json:"password"`
}

const password = "super_secret"

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	apiKey := "my-secret-key"
	name := p.ByName("name")
	payload := Payload{
		Message:  "Hello " + name,
		Password: apiKey,
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func main() {
	router := httprouter.New()
	router.POST("/hello/:name", hello)

	err:= http.ListenAndServe("0.0.0.0:5001", router)
	if err!= nil{}
}
