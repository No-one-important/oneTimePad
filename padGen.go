package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	length, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < length; i++ {
		fmt.Printf("%v\n", rand.Intn(26))
	}
}
