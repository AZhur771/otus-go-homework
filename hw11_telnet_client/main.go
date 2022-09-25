package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10, "timeout specified with time unit")
}

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintln(os.Stderr, "Please provide host and port arguments")
		os.Exit(1)
	}

	address := net.JoinHostPort(flag.Arg(0), flag.Arg(1))
	telnetClient := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	if err := telnetClient.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot connect to %s", address)
		return
	}

	fmt.Fprintf(os.Stderr, "...Connected to %s\n", address)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(
		signalCh,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	defer close(signalCh)

	done := make(chan struct{})
	defer close(done)

	go func() {
		for {
			err := telnetClient.Send()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				done <- struct{}{}
				return
			}
		}
	}()

	go func() {
		for {
			err := telnetClient.Receive()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				done <- struct{}{}
				return
			}
		}
	}()

	go func() {
		<-signalCh
		done <- struct{}{}
	}()

	<-done
	telnetClient.Close()
}
