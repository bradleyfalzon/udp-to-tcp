package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	flagListen := flag.String("listen", "", "host:port of UDP address to listen on, use :port for all interfaces")
	flagBackend := flag.String("backend", "", "host:port of remote tcp connection")
	flagTimeout := flag.Int("timeout", 5, "tcp connection timeout")
	flag.Parse()

	if *flagBackend == "" || *flagListen == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Connect to remote TCP.
	be, err := net.DialTimeout("tcp", *flagBackend, time.Duration(*flagTimeout)*time.Second)
	if err != nil {
		log.Fatalf("Could not connect to backend: %v", err)
	}

	// Listen to UDP.
	addr, err := net.ResolveUDPAddr("udp", *flagListen)
	if err != nil {
		log.Fatalf("Could not resolve udp: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Could not listen on udp: %v", err)
	}
	defer conn.Close()

	log.Printf("Listening for UDP on %q sending to TCP %q", *flagListen, *flagBackend)

	// Receive packets
	buf := make([]byte, 1500)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading from udp on server listen loop: ", err)
			continue
		}
		_, err = be.Write(buf[:n])
		if err != nil {
			log.Fatalf("Could not write packet to tcp connection: %v", err)
		}

	}
}
