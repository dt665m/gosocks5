package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dt665m/gosocks5"
)

func main() {
	go func() {
		fmt.Println("Starting Test Server 8888")
		mux := http.NewServeMux()
		mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			fmt.Fprintf(w, "your ip is %v", ip)
		})
		http.ListenAndServe(":8888", mux)
	}()
	conf := &socks5.Config{}
	conf.Credentials = socks5.StaticCredentials(map[string]string{"hello": "world"})
	conf.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8080
	fmt.Println("Starting SOCKS5 on 8080")
	if err := server.ListenAndServe("tcp", ":8080"); err != nil {
		panic(err)
	}
}
