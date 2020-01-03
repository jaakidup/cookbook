package bytestrings

import "bytes"

import "io"

import "io/ioutil"

// Buffer demonstrates some tricks for initializing bytes buffers
// These buffer implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	// we'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)

	// there are a number of ways to create a buffer from
	// the raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	// alternatively
	b = bytes.NewBuffer(rawBytes)
	// and avoiding the initial bytes altogether
	b = bytes.NewBufferString(rawString)
	return b
}

// toString is an example of taking an io.Reader and consuming
// it all, then returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
