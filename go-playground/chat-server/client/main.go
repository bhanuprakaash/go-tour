package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("\nDisconnected from server.")
			os.Exit(0)
		}
	}()

	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
