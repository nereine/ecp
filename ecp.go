package main

import (
	"os"
)

func readF(fname string) []byte {
	b, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return b
}

func writeF(src string, dst string) {
	data := readF(src)
	err := os.WriteFile(dst, data, 0777)
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(444)
	}
	src := os.Args[1]
	dst := os.Args[2]
	writeF(src, dst)
}
