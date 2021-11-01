package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	log.Println("My addresses:")
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			log.Printf("\tIP on %s is %s", iface.Name, addr)
		}
	}

	tick := time.NewTicker(100 * time.Millisecond)
	for range tick.C {
		resp, err := http.Get("https://www.google.com/")
		if err != nil {
			log.Printf("Error getting www.google.com: %v", err)
			continue
		} else if resp.StatusCode != http.StatusOK {
			log.Printf("Non-200 return from www.google.com: %d", resp.StatusCode)
			continue
		}
		break
	}
	tick.Stop()

	log.Printf("Network is working now, sleeping forever...")
	for {
		time.Sleep(time.Hour)
	}
}
