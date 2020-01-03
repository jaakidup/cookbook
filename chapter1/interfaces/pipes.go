package interfaces

import "io"

import "os"

// PipeExample gives some more examples of using interfaces
func PipeExample() error {
	// the pipe reader and pipe writer implement
	// io.Reader and io.Writer
	r, w := io.Pipe()

	// this needs to be run in a seperate go routine
	// as it will block waiting for the reader
	// close at the end for cleanup
	go func() {
		// for now we'll write someting basic
		// this could also be used to encode json
		// base64 encode etc.
		w.Write([]byte("test"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
