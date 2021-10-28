package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danmrichards/container-runner/internal/clients"
	"github.com/danmrichards/container-runner/internal/server"
)

var host, port string

func main() {
	flag.StringVar(&host, "host", "0.0.0.0", "host for the HTTP clients")
	flag.StringVar(&port, "port", "8080", "port for the HTTP clients")
	flag.Parse()

	m := clients.NewManager()

	go logClients(m)

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	svr, err := server.NewServer(port, host, m)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP server listening on " + port)
	go func() {
		if err = svr.Serve(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sigChan

	log.Println("shutting down")
	svr.Close()
}

func logClients(mgr *clients.Manager) {
	for {
		<-time.After(1 * time.Second)

		log.Println(mgr)
	}
}
