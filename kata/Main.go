package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(DNAtoRNA("GATU"))

	fmt.Println(Race(10, 20, 333))

	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
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

func Race(v1, v2, g int) [3]int {

	distancePerHour := (v2 - v1)

	timeInSeconds := g * 3600 / distancePerHour

	// Could use the time package if return of an array is not required.
	// https://golang.org/pkg/time/#Duration
	fmt.Println(time.Duration(timeInSeconds) * time.Second)

	return [3]int{timeInSeconds / 3600, (timeInSeconds / 60) % 60, timeInSeconds % 60}
}

func PrinterError(s string) string {
	r := []rune(s)
	// get the last letter of the string
	// endLetter := r[len(r)-1]

	// this question is actually easier, because the maximum is define, as m
	end := rune('m')

	length := len(r)

	var count int

	for _, v := range r {
		if v > end {
			count++
		}
	}

	return fmt.Sprintf("%d/%d", count, length)
}
