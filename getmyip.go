package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var ipv4 bool
	var ipv6 bool
	var verbose bool

	flag.BoolVar(&ipv4, "4", false, "Get your external IP v4 address")
	flag.BoolVar(&ipv6, "6", false, "Get your external IP v6 address")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Usage = func() {
		fmt.Println("Usage: getmyip [options]")
		fmt.Println("     -4 Get your external IPv4 address")
		fmt.Println("     -6 Get your external IPv6 address")
		fmt.Println("     -v Verbose output")
	}

	flag.Parse()

	if ipv4 {
		output("Your IPv4 address:", get("https://ip4.5ec.nl"), verbose)
	}

	if ipv6 {
		output("Your IPv6 address:", get("https://ip6.5ec.nl"), verbose)
	}

	if !ipv4 && !ipv6 {
		output("Your IP address:", get("https://ip.5ec.nl"), verbose)
	}
}

func output(title string, ip string, verbose bool) {
	if verbose {
		fmt.Print(title, " ", ip)
	} else {
		fmt.Print(ip)
	}
}

func get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		print(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		print(err)
		os.Exit(1)
	}

	return string(body)
}
