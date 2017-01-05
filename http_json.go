package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PayLoad struct {
	Content string
}

func HttpJson() {
	fmt.Println("Http JSON")
	// create the pipe
	pr, pw := io.Pipe()

	go func(pw *io.PipeWriter) {
		// close the writer, so the reader knows there's no more data
		defer pw.Close()

		// write json data into the pipereader through the pipewriter
		if err := json.NewEncoder(pw).Encode(&PayLoad{Content: "Hello Pipe!"}); err != nil {
			log.Fatal(err)
		}
	}(pw)

	// the json from the pipewriter lands in the pipereader
	// and we send it off...
	if _, err := http.Post("http://example.com", "application/json", pr); err != nil {
		log.Fatal(err)
	}
}
