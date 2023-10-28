package main

import (
	"github.com/Peter-Roh/gocoin/explorer"
	"github.com/Peter-Roh/gocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
