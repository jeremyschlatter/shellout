// Package shellout provides a convenince wrapper for os/exec.
package shellout

import (
	"io"
	"os/exec"
)

type reader struct {
	cmd *exec.Cmd
	r   io.ReadCloser
}

func (r *reader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	if err == io.EOF {
		r.cmd.Wait()
	}
	return n, err
}

// Start runs an external command which reads from stdin and writes to stdout.
// This library assumes the caller will read stdout until EOF. If not, resources
// (in particular file descriptors) may leak.
func Start(stdin io.Reader, name string, arg ...string) (stdout io.Reader, err error) {
	cmd := exec.Command(name, arg...)
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	go func() {
		io.Copy(inPipe, stdin)
		inPipe.Close()
	}()
	return &reader{cmd, outPipe}, nil
}
