package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/FryDay/go-electrumx/electrumx"
)

var (
	tlsFlag = flag.Bool("tls", true, "dial with TLS")
	timeout = flag.Duration("timeout", time.Second*3, "ping timeout")
	debug   = flag.Bool("debug", false, "turn on debug output")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "%s <target>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := realMain(); err != nil {
		log.Fatal(err)
	}

}
func realMain() error {

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if *debug {
		electrumx.DebugMode = true
	}

	start := time.Now()
	node := electrumx.NewNode()

	go func() {
		if err := <-node.Errors(); err != nil {
			log.Fatalf("draining node errors: %s", err)
		}
	}()
	var config *tls.Config
	if *tlsFlag {
		config = &tls.Config{}
	}
	if err := node.Connect(ctx, flag.Arg(0), config); err != nil {
		return fmt.Errorf("connect to %s: %w", flag.Arg(0), err)
	}

	if err := node.Ping(ctx); err != nil {
		return fmt.Errorf("ping %s: %w", flag.Arg(0), err)
	}

	log.Printf("pinged %s in %s", flag.Arg(0), time.Since(start))
	return nil
}
