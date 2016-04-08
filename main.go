package main

import (
	"github.com/mougeli/go-socks5"
	"os"
)

func main() {

	// Create a SOCKS5 server
	conf := &socks5.Config{
	}
	if (os.Getenv("AUTH_USERNAME") != "") {
		c := socks5.StaticCredentials{}
		c[os.Getenv("AUTH_USERNAME")] = os.Getenv("AUTH_PASSWORD")
		conf = &socks5.Config{
			Credentials:c,
		}
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", os.Getenv("LISTEN")); err != nil {
		panic(err)
	}
}
