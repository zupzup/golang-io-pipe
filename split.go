package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func Split() {
	fmt.Println("Split")
	// set up pipe
	pr, pw := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(2)

	// read file
	f, err := os.Open("./fruit.txt")
	if err != nil {
		log.Fatal(err)
	}

	// set up teereader
	tr := io.TeeReader(f, pw)

	go func() {
		defer wg.Done()
		defer pw.Close()

		fmt.Println("Starting to send...")
		// teereader has the data from the file and also writes it to the pipereader
		_, err := http.Post("https://example.com", "text/html", tr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Sending successful.")
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Starting to print...")
		// read from the pipereader, which gets written to via the pipewriter, to stdout
		if _, err := io.Copy(os.Stdout, pr); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Printing successful.")
	}()

	wg.Wait()
}
