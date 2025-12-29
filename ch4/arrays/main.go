package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"math/bits"
	"os"
)

func main() {
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	var a [3]int
	// var q [3]int = [3]int{1, 2}
	q := [...]int{1, 2, 3}

	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "S"}
	r := [...]int{99: 0}

	for i, e := range a {
		fmt.Printf("%d, %d\n", i, e)
	}

	fmt.Println(q[2])
	fmt.Println(symbol[EUR])
	fmt.Println(len(r))

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x \n %x \n %t \n %T\n", c1, c2, c1 == c2, c1)
	fmt.Println(countBits(c1, c2))

	mutateArray(&a)
	fmt.Println(a)

	sha384Flag := flag.Bool("sha384", false, "print SHA384 hash")
	sha512Flag := flag.Bool("sha512", false, "print SHA512 hash")
	flag.Parse()

	// Read from standard input
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}

	printShaValues(input, *sha384Flag, *sha512Flag)
}

func mutateArray(a *[3]int) {
	a[2] = 32
}

// 4.1 Write a function that counts the number of bits that are different in two SHA256 hashes.
func countBits(sha1, sha2 [32]uint8) int {
	count := 0

	for i := 0; i < 32; i++ {
		diff := sha1[i] ^ sha2[i]
		count += bits.OnesCount8(diff)
	}

	return count
}

// Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.
func printShaValues(message []byte, showSha384, showSha512 bool) {
	switch {
	case showSha384:
		hash := sha512.Sum384(message)
		fmt.Printf("SHA384: %x\n", hash)

	case showSha512:
		hash := sha512.Sum512(message)
		fmt.Printf("SHA512: %x\n", hash)

	default:
		hash := sha256.Sum256(message)
		fmt.Printf("SHA256: %x\n", hash)
	}
}
