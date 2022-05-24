package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// TODO: set a http client timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Close the response body on the return.
	defer resp.Body.Close()

	// Write the response to stdout.
	io.Copy(os.Stdout, resp.Body)
}
