package main

import (
	"encoding/json"
	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/middlewaregruppen/kubez/pkg/info"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/cgroup", func(w http.ResponseWriter, r *http.Request) {
		rt := info.GetCgroup()
		b, _ := json.Marshal(rt)
		w.Write(b)
	})

	box := packr.NewBox("../../web/frontend/dist")
	r.Handle("/", http.FileServer(box))

	http.ListenAndServe(":3000", r)
}
