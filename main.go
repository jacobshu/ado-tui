package main

import (
	"fmt"
	"os"

	"github.com/jacobshu/ado-tui/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("error executing command\n")
		os.Exit(1)
	}
}
