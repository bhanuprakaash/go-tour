package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

type Clock struct {
	Name    string
	Host    string
	Scanner *bufio.Scanner
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("clockall NewYork=localhost:8010...")
		return
	}
	var clocks []*Clock
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			log.Printf("Bad argument format: %s. Use City=Host:Port", arg)
			continue
		}

		conn, err := net.Dial("tcp", parts[1])
		if err != nil {
			log.Fatal("something went wrong", err)
		}
		defer conn.Close()
		clocks = append(clocks, &Clock{
			Name:    parts[0],
			Host:    parts[1],
			Scanner: bufio.NewScanner(conn),
		})
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	for _, c := range clocks {
		fmt.Fprintf(tw, "%s\t", c.Name)
	}
	fmt.Fprintf(tw, "\n")
	for _, c := range clocks {
		fmt.Fprintf(tw, "%s\t", strings.Repeat("-", len(c.Name)))
	}
	fmt.Fprintf(tw, "\n")
	tw.Flush()

	for {
		for _, c := range clocks {
			if c.Scanner.Scan() {
				fmt.Fprintf(tw, "%s\t", c.Scanner.Text())
			} else {
				fmt.Fprintf(tw, "OFFLINE\t")
			}
		}
		fmt.Fprintf(tw, "\n")
		tw.Flush()
	}
}
