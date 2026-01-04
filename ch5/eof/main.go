package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample_file.txt")

	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			fmt.Printf("Read chunk: %q\n", buffer[:n])
		}

		if err == io.EOF {
			fmt.Println("\nResult: End of File (EOF) reached. Stopping.")
			break
		}

		if err != nil {
			fmt.Println("Unexpected error:", err)
			break
		}
	}
}
