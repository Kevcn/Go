package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DNAtoRNA("GATU"))
}

func DNAtoRNA(dna string) string {
	var sb strings.Builder

	chars := strings.Split(dna, "")

	for _, v := range chars {
		if v == "T" {
			sb.WriteString("U")
		} else {
			sb.WriteString(v)
		}
	}

	return sb.String()

	// Best practice - neat 1 liner.... its all about knowing your library
	// return strings.Replace(dna, "T", "U", -1)
}
