package daemon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"4discovery/store"
)

func (d *Daemon) processHttpReq(c chan<- error) {
	handler := http.NewServeMux()

	handler.HandleFunc("/status", status)
	handler.HandleFunc("/", d.index)
	handler.HandleFunc("/services", d.serviceHandler)

	c <- http.ListenAndServe(d.HTTPEndpoint, handler)
}

func (d *Daemon) index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/services", http.StatusMovedPermanently)
}

func (d *Daemon) serviceHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string][]store.Service)

	response["services"] = d.store.List()
	jsonResp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonResp))
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "4discovery lives!")
}
