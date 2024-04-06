package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCmd(cmd string, args ...string) (stdout string, stderr string, err error) {
	c := exec.Command(cmd, args...)
	stdoutbuf, stderrbuf := new(bytes.Buffer), new(bytes.Buffer)
	c.Stdout = stdoutbuf
	c.Stderr = stderrbuf
	err = c.Run()
	stdout = string(stdoutbuf.Bytes())
	stderr = string(stderrbuf.Bytes())
	if err != nil {
		return stdout, stderr, fmt.Errorf("%s: %s", err.Error(), stderr)
	}
	return
}
