package main

import (
	"fmt"
)

type DatabaseConfig struct {
	ConnectionString string
}

type WebServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Database  DatabaseConfig
	WebServer WebServerConfig
}

func main() {
	dc := DatabaseConfig{ConnectionString: "GoLang is super cool!"}
	wsc := WebServerConfig{
		Host: "localhost",
		Port: 8080, // end with a comma!
	}
	c := Config{
		Database:  dc,
		WebServer: wsc,
	}

	p := &c

	fmt.Println(dc.ConnectionString) // GoLang is super cool!
	fmt.Println(wsc.Host)            // localhost
	fmt.Println(p.WebServer)         // {localhost 8080}
	fmt.Println((*p).WebServer)      // same as p.Host
}
