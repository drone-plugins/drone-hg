package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

type Plugin struct {
	Repo  Repo
	Build Build
	Netrc Netrc
}

func (p Plugin) Exec() error {
	if p.Build.Path != "" {
		err := os.MkdirAll(p.Build.Path, 0777)
		if err != nil {
			return err
		}
	}

	err := writeHgrc(p.Netrc.Machine, p.Netrc.Login, p.Netrc.Password)
	if err != nil {
		return err
	}

	var cmds []*exec.Cmd

	if isDirEmpty(filepath.Join(p.Build.Path, ".hg")) {
		cmds = append(cmds, initHg())
	}
	cmds = append(cmds, pull(p.Build.Commit, p.Repo.Clone))
	cmds = append(cmds, update(p.Build.Commit))

	for _, cmd := range cmds {
		cmd.Dir = p.Build.Path
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

func initHg() *exec.Cmd {
	return exec.Command(
		"hg",
		"init",
		".",
	)
}

func pull(rev, path string) *exec.Cmd {
	return exec.Command(
		"hg",
		"pull",
		"--rev",
		rev,
		path,
	)
}

func update(rev string) *exec.Cmd {
	return exec.Command(
		"hg",
		"update",
		rev,
	)
}
