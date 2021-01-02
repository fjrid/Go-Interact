package interact

import (
	"bytes"
	"io"
	"os/exec"
)

/*
	CmdInteract used for persisting config to run the exec.Command().
	Other than that, this struct can be used for any configuration
	for this package.
*/
type CmdInteract struct {
	command string
	args    []string
	exec    *exec.Cmd

	Silent bool

	StdOut []byte
	StdErr []byte
}

/*
	This struct implementing io.Writer to storing current bytes
*/
type capturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

/*
	Create new instance for capturingPassThroughWriter struct
*/
func newCapturingPassThroughWriter(w io.Writer) *capturingPassThroughWriter {
	return &capturingPassThroughWriter{
		w: w,
	}
}

/*
	Write method for writing buffer
*/
func (w *capturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

/*
	Bytes method for converting to the bytes
*/
func (w *capturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}
