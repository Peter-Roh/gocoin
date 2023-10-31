package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Peter-Roh/gocoin/explorer"
	"github.com/Peter-Roh/gocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to gocoin!\n\n")
	fmt.Printf("Please use the following flags...\n\n")
	fmt.Printf("-port:		Set the port of the server\n")
	fmt.Printf("-mode:   	Choose between 'html' and 'rest'\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Set the port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
