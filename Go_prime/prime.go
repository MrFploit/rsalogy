package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	q, _ := strconv.Atoi(input)
	p, _ := rand.Prime(rand.Reader, q)
	fmt.Println(p)
}
