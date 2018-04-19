package main

import (
	"log"
	"os"

	"github.com/dt665m/gosocks5"
)

func main() {
	conf := &socks5.Config{}
	conf.Credentials = socks5.StaticCredentials(map[string]string{"hello": "world"})
	conf.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", ":8080"); err != nil {
		panic(err)
	}
}
