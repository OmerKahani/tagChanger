package main

import (
	"os"

	"github.com/OmerKahani/tagChanger/cmd"
)

func main() {
	if err := tagChanger.GetCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
