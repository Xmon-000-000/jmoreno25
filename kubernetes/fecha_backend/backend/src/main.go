package main

import (
    "net/http"
	"os"
	"time"
	"encoding/json"
)

type HansOn struct {
	Time time.Time `json:"time"`
	Hostname string `json:"hostname"`
}

func ServerHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HansOn{
		Time: time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error: No se ha podido procesar la request"))
		return
	}

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResp)
}

func main() {
    http.HandleFunc("/", ServerHTTP)
    http.ListenAndServe(":9090", nil)
}
