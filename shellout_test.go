package shellout

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Example() {
	stdout, _ := Start(strings.NewReader("example"), "cat")
	io.Copy(os.Stdout, stdout)
	// Output:
	// example
}

func TestBadFileDescriptor(t *testing.T) {
	stdout, err := Start(strings.NewReader("example"), "cat")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
}
