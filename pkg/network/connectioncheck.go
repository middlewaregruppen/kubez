package network

import (
	"encoding/json"
	"fmt"
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
		res.Error = fmt.Sprintf("%s", err.Error())
		res.ErrorType = reflect.TypeOf(err).String()

	} else {
		res.Successful = true
		res.Address = conn.RemoteAddr().String()
		conn.Close()
	}

	b, _ := json.Marshal(res)

	w.Write(b)

}
