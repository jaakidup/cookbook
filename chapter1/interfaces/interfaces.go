package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// NetWriter writes a log over the network
// fits the io.Writer interface
type NetWriter struct {
	URL string
}

func (nw *NetWriter) Write(p []byte) (n int, err error) {
	sender := bytes.NewReader(p)
	if _, err = http.Post(nw.URL, "application/text", sender); err != nil {
		return 0, err
	}
	return len(p), err
}

// Copy copies data from in to out, first directly,
// then using a buffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {

	// let's start a logging server that just prints out the []bytes to console
	// in a seperate goroutine
	go func() {
		http.HandleFunc("/logger", func(w http.ResponseWriter, r *http.Request) {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("http://localhost:9000/logger", string(bytes))
		})
		log.Println("Starting server on :9000")
		log.Fatalln(http.ListenAndServe(":9000", nil))
	}()

	netwriter := &NetWriter{
		URL: "http://localhost:9000/logger",
	}

	// write to out and to stdout as well as to netwriter (network log writer)
	w := io.MultiWriter(out, os.Stdout, netwriter)
	// a standard copy, this can be dangerous if there's a lot of data in in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	in.Seek(0, 0)

	// buffered write using 64 byte chunks
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// let's print a new line
	fmt.Println()
	return nil
}
