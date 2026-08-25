package network

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
)

// ConnectionStatus ...
type ConnectionStatus struct {
	Successful bool   `json:"success"`
	Address    string `json:"address"`
	Error      string `json:"error"`
	ErrorType  string `json:"errorType"`
}

// HandleCheckConnection ...
func HandleCheckConnection(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := ConnectionStatus{}
	conn, err := net.DialTimeout("tcp", vars["target"], time.Second*10)

	if err != nil {
		res.Successful = false
		res.Error = err.Error()
		res.ErrorType = reflect.TypeOf(err).String()

	} else {
		res.Successful = true
		res.Address = conn.RemoteAddr().String()
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection to %s: %s", vars["target"], err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to write connection check response: %s", err)
	}

}
