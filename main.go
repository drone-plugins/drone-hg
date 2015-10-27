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

	"github.com/drone/drone-plugin-go/plugin"
)

var hgrcFile = `
[auth]
bb.prefix = https://bitbucket.org
bb.username = %s
bb.password = %s
`

// Params stores the git clone parameters used to
// configure and customzie the git clone behavior.
type Params struct {
}

func main() {
	v := new(Params)
	r := new(plugin.Repo)
	b := new(plugin.Build)
	w := new(plugin.Workspace)
	plugin.Param("repo", r)
	plugin.Param("build", b)
	plugin.Param("workspace", w)
	plugin.Param("vargs", &v)
	plugin.MustParse()

	err := run(r, b, w, v)
	if err != nil {
		os.Exit(1)
	}
}

// run clones the repository and build revision
// into the build workspace.
func run(r *plugin.Repo, b *plugin.Build, w *plugin.Workspace, v *Params) error {

	err := os.MkdirAll(w.Path, 0777)
	if err != nil {
		fmt.Printf("Error creating directory %s. %s\n", w.Path, err)
		return err
	}

	// generate the .hgrc file
	if err := writeHgrc(w); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	var cmds []*exec.Cmd
	cmds = append(cmds, clone(b))
	cmds = append(cmds, update(b))

	for _, cmd := range cmds {
		cmd.Dir = w.Path
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		trace(cmd)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func clone(b *plugin.Build) *exec.Cmd {
	return exec.Command(
		"hg",
		"clone",
		"--branch",
		b.Branch,
	)
}

func update(b *plugin.Build) *exec.Cmd {
	return exec.Command(
		"hg",
		"update",
		b.Commit,
	)
}

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}

// Writes the hgrc file.
func writeHgrc(in *plugin.Workspace) error {
	if in.Netrc == nil || len(in.Netrc.Machine) == 0 {
		return nil
	}
	out := fmt.Sprintf(
		hgrcFile,
		in.Netrc.Machine, // TODO this may require adding http(s) prefix
		in.Netrc.Login,
		in.Netrc.Password,
	)
	home := "/root/.hg"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}
	path := filepath.Join(home, "hgrc")
	return ioutil.WriteFile(path, []byte(out), 0600)
}

func isDirEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return true
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false
}
