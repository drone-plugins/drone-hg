package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

const hgrcFile = `
[auth]
drone.prefix = %s
drone.username = %s
drone.password %s
drone.schemes = http https
`

// helper function to write a hgrc file.
func writeHgrc(machine, login, password string) error {
	if machine == "" {
		return nil
	}
	out := fmt.Sprintf(
		hgrcFile,
		machine,
		login,
		password,
	)

	home := "/root"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}
	path := filepath.Join(home, ".netrc")
	return ioutil.WriteFile(path, []byte(out), 0600)
}

// helper function returns true if directory dir is empty.
func isDirEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return true
	}
	defer f.Close()

	_, err = f.Readdir(1)
	return err == io.EOF
}

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
