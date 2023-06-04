package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	m := make(map[string]int)

	data, err := os.ReadFile("proverb.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	for ln, line := range strings.Split(s, "\n") {
		n := ln + 1
		for _, c := range strings.Split(line, "") {
			m[c] = strings.Count(line, c)
		}
		fmt.Printf("%d. %s\n", n, line)
		for k, v := range m {
			fmt.Printf("'%s'= %d, ", k, v)
		}
		fmt.Printf("\n\n")
	}
}
