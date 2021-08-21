package main

import (
	"flag"
	"log"
	"os"

	"github.com/nasum/hatebu-crawler/lib"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	switch os.Args[1] {
	case "bookmark":
		cmd := flag.NewFlagSet("bookmark", flag.ExitOnError)
		target := cmd.String("target", "", "crawl target")
		err := cmd.Parse(os.Args[2:])

		if err != nil {
			log.Fatal(err)
		}

		err = lib.GetEntries(*target)

		if err != nil {
			log.Fatal(err)
		}
	case "top":
		err := lib.GetTop()
		if err != nil {
			log.Fatal(err)
		}
	}
}
