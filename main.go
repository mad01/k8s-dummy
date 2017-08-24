package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	hostname     string
	httpAddr     string
	buildVersion string
)

func main() {
	flag.StringVar(&httpAddr, "http", "127.0.0.1:80", "HTTP service address")
	flag.Parse()

	var err error
	hostname, err = os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	errChan := make(chan error, 10)

	router := mux.NewRouter()
	router.HandleFunc("/version", httpVersionHandler)
	router.HandleFunc("/healthz", httpHealthzHandler)
	router.HandleFunc("/", httpHostnameHandler)

	http.Handle("/", router)
	go func() {
		errChan <- http.ListenAndServe(httpAddr, nil)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				fmt.Printf("%s - %s\n", hostname, err)
				os.Exit(1)
			}
		case <-signalChan:
			fmt.Printf("%s - Shutdown signal received, exiting...\n", hostname)
			os.Exit(0)
		}
	}
}
