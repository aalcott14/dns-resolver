package server

import (
	"io"
	"log"

	"github.com/aalcott14/dns-resolver/resolver"
)

func NewServer() {
	dnsClient := resolver.NewDnsClient()

	resp, err := dnsClient.Get("https://www.twitter.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
