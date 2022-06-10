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
	var output []byte

	for i := 0; i < len(text); i++ {
		output = append(output, caesar(text[i], pad[i]))
	}
	return string(output)
}

func caesar(r byte, shift int) byte {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	return r + byte(shift)
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

	fmt.Printf("%v", encrypt(strings.ToLower(string(text)), padNum))
}
