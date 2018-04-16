package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("main-engine")

var (
	cp app.ConfigProvider
)

func main() {

	if cp == nil {
		// Use default config provider
		cp = app.DefaultConfigProvider()
	}

	app, err := cp.GetApp()
	if err != nil {
        	fmt.Println(err.Error())
        	os.Exit(1)
    }
    
    e, err := engine.New(app)
	if err != nil {
		log.Errorf("Failed to create engine instance due to error: %s", err.Error())
		os.Exit(1)
	}

	err = e.Start()
	if err != nil {
		log.Errorf("Failed to start engine due to error: %s", err.Error())
		os.Exit(1)
	}

	exitChan := setupSignalHandling()

	code := <-exitChan

	e.Stop()

	os.Exit(code)
}

func setupSignalHandling() chan int {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGHUP
			case syscall.SIGHUP:
				exitChan <- 0
			// kill -SIGINT/Ctrl+c
			case syscall.SIGINT:
				exitChan <- 0
			// kill -SIGTERM
			case syscall.SIGTERM:
				exitChan <- 0
			// kill -SIGQUIT
			case syscall.SIGQUIT:
				exitChan <- 0
			default:
				logger.Debug("Unknown signal.")
				exitChan <- 1
			}
		}
	}()

	return exitChan
}
