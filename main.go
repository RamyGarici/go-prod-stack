package main

import (
	"log"
	"net/http"
	"flag"
	"os"
	"time"
)

func main() {
	check := flag.Bool("healthcheck", false, "Check API's health")
	flag.Parse()
	if *check{
		client := &http.Client{Timeout: 2 * time.Second}
		resp,err := client.Get("http://127.0.0.1:8080/health")
		if err!=nil||resp.StatusCode != http.StatusOK{
			os.Exit(1)
		}
		resp.Body.Close()
		os.Exit(0)
	}
	
	s := http.NewServeMux()
	s.HandleFunc("GET /health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
		log.Println("ok")
	})
	
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatal(err)
	}
    
	



}
