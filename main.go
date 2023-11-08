package main

import (
	"github.com/Peter-Roh/gocoin/cli"
	"github.com/Peter-Roh/gocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
