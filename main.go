package interact

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

/*
	Initiate method for initialize CmdInteract, this method can be used for
	initialize default configuration in package
*/
func Initiate(command string, args ...string) *CmdInteract {
	cmdInteract := CmdInteract{
		Silent: false,
	}

	cmdInteract.command = command
	cmdInteract.args = args
	cmdInteract.exec = exec.Command(
		cmdInteract.command,
		cmdInteract.args...,
	)

	return &cmdInteract
}

/*
	Run method for run the command configration
*/
func (cmi *CmdInteract) Run() error {
	cmd := cmi.exec

	var errStdout, errStderr error

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var stdout, stderr *capturingPassThroughWriter

	if cmi.Silent {
		stdout = newCapturingPassThroughWriter(&bytes.Buffer{})
		stderr = newCapturingPassThroughWriter(&bytes.Buffer{})
		cmd.Stdin = stdoutIn
	} else {
		stdout = newCapturingPassThroughWriter(os.Stdout)
		stderr = newCapturingPassThroughWriter(os.Stderr)
		cmd.Stdin = os.Stdin
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("cmd.Start() failed with '%s'", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()

	if err != nil {
		return fmt.Errorf("cmd.Run() failed with %s", err)
	}
	if errStdout != nil || errStderr != nil {
		return fmt.Errorf("failed to capture stdout or stderr")
	}

	cmi.StdOut = stdout.Bytes()
	cmi.StdErr = stderr.Bytes()

	return nil
}

/*
Convert command to string
*/
func (cmi *CmdInteract) String() string {
	return cmi.exec.String()
}
