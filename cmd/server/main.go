package main

import (
	"flag"
	"fmt"
	"github.com/laybatin/bodymirror-app-backend/config"
	backend "github.com/laybatin/bodymirror-app-backend/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	backend.Init()

}
