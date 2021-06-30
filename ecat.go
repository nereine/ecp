package main

import (
	"os"
)

func read(fname string) []byte {
	b, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return b
}

func cat(src string) {
	os.Stdout.Write(read(src))
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(69)
	}

	files := os.Args[1:]
	for _, src := range files {
		cat(src)
	}
}
