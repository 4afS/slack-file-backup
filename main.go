package main

import (
	"flag"
)

func main() {
	var (
		channelID string
		nfiles    int
		filetype  string
	)

	flag.StringVar(&channelID, "c", "", "select channel ID (default all)")
	flag.IntVar(&nfiles, "n", 100, "get n file names")
	flag.StringVar(&filetype, "t", "", "download only selected file type (default all)")

	flag.Parse()
}
