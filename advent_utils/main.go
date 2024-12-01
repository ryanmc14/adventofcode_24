package advent_utils

import (
	"bufio"
	"log"
	"os"
)

func Input(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer f.Close()
	return bufio.NewScanner(f)
}
