package main

import (
	"log"
	"main/handler"
	"net/http"
	"os"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/SampleHttpTrigger", handler.SampleHttpTrigger)

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
