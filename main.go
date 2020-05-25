package main

import (
	"os"
	"schuck017/go-distributed-blockchain-database/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
