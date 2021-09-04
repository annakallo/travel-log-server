package main

import (
	"encoding/json"
	"fmt"
	"github.com/annakallo/travel-log-server/server"
	"net/http"
)

func main() {

	r := server.NewRouter()
	e := http.ListenAndServe(":12345", r)
	// @TODO needs to be tested this error handling
	responseJson, _ := json.Marshal(e)
	if e != nil {
		fmt.Printf("Error in listen and serve %s", string(responseJson))
	}
}
