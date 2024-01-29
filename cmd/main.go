package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.comleobaiano//sniffer/pkg/capture"
)

func main() {
	stopCapture := make(chan struct{})
	go func() {
		err := capture.StartCapture("wlan0")
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