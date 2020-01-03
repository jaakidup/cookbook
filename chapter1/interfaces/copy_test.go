package interfaces

import "testing"

import "bytes"

// import "github.com/jaakidup/go-cookbook/chapter1/interfaces"

func TestCopy(t *testing.T) {

	data := []byte("something")
	in := bytes.NewReader(data)
	out := &bytes.Buffer{}

	if err := Copy(in, out); err != nil {
		t.Error("Failed Copy")
	}
	// needs to be data+data because copy copies twice
	if string(data)+string(data) != out.String() {
		t.Error("expected data to be the same on in and out")
		t.Error("Expected ", string(data), " got ", out)
	}

}
