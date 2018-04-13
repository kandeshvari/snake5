package main

import (
	"github.com/getlantern/go-socks5"
	"flag"
)

var address string
var proto string
var username string
var password string
func init() {
	flag.StringVar(&address, "address", "0.0.0.0:1081", "Address to bind")
	flag.StringVar(&proto, "proto", "tcp", "Use proto (tcp|udp?)")
	flag.StringVar(&username, "username", "", "Username for auth")
	flag.StringVar(&password, "password", "", "Password for auth")
	flag.Parse()
}


func main() {
	var conf *socks5.Config
	if username != "" && password != "" {
		conf = &socks5.Config{
			AuthMethods: []socks5.Authenticator{
				socks5.UserPassAuthenticator{
					Credentials: &AuthStore{
						Username: username,
						Password: password,
					},
				},
			},
		}
	} else {
		conf = &socks5.Config{}
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe(proto, address); err != nil {
		panic(err)
	}
}
