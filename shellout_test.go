package shellout

import (
	"io"
	"os"
	"strings"
)

func ExampleShellout() {
	stdout, _ := Start(strings.NewReader("example"), "cat")
	io.Copy(os.Stdout, stdout)
	// Output:
	// example
}
