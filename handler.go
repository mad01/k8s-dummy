package main

import (
	"fmt"
	"net/http"
)

func httpHostnameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server hostname: %s", hostname)
}

func httpHealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Healthz-Header", "Running")
	fmt.Fprint(w, "ok")
}
func httpVersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server Version: %s", buildVersion)
}
