package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	srv:= &http.Server{Addr: ":8080", Handler: s}
	go func ()  {
		if err:= srv.ListenAndServe(); err!=nil && !errors.Is(err,http.ErrServerClosed){
			log.Fatal(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)

    
	



}
