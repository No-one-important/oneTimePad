package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func encrypt(text string, pad []int) string {
	r := []rune(text)
	var output []rune

	for i, j := range r {
		output = append(output, caesar(j, pad[i]*-1))
	}
	return string(output)
}

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func main() {
	padFile := os.Args[1]
	textFile := os.Args[2]

	text, err := ioutil.ReadFile(textFile)
	if err != nil {
		log.Fatal(err)
	}

	pad, err := ioutil.ReadFile(padFile)
	if err != nil {
		log.Fatal(err)
	}

	padSplit := strings.Split(string(pad), "\n")
	var padNum []int
	var num int

	for _, s := range padSplit {
		num, _ = strconv.Atoi(s)
		padNum = append(padNum, num)
	}

	fmt.Printf("%v", strings.ReplaceAll(strings.ReplaceAll(encrypt(string(text), padNum), "T", " "), "A>", "\n"))
}
