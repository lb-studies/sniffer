package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lb-studies/sniffer/pkg/capture"
)

func main() {
	stopCapture := make(chan struct{})
	go func() {
		err := capture.StartCapture("wlp0s20f3")
		if err != nil {
			log.Fatal(err)
		}
		close(stopCapture)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	// Encerre a captura
	fmt.Println("Ending the capture...")
	close(stopCapture)
}