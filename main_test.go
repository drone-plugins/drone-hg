package main

import (
	"github.com/drone/drone-plugin-go/plugin"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var commits = []struct {
	path   string
	clone  string
	event  string
	branch string
	commit string
	file   string
	data   string
}{
	// first commit
	{
		path:   "cedk/drone-hg-test",
		clone:  "https://bitbucket.org/cedk/drone-hg-test",
		event:  plugin.EventPush,
		branch: "default",
		commit: "37526193d0139f188b20e5c8bed8fc0640c38627",
		file:   "test",
		data:   "",
	},
	// head commit
	{
		path:   "cedk/drone-hg-test",
		clone:  "https://bitbucket.org/cedk/drone-hg-test",
		event:  plugin.EventPush,
		branch: "default",
		commit: "6e4637b5ef305536115bd05e97c913fc7cdcc69d",
		file:   "test",
		data:   "Test\n",
	},
}

func TestClone(t *testing.T) {
	for _, c := range commits {
		dir := setup()

		r := &plugin.Repo{Clone: c.clone}
		b := &plugin.Build{Commit: c.commit, Branch: c.branch, Event: c.event}
		w := &plugin.Workspace{Path: dir}
		v := &Params{}

		if err := run(r, b, w, v); err != nil {
			t.Errorf("Expected successful clone. Got error. %s.", err)
		}

		data := readFile(dir, c.file)
		if data != c.data {
			t.Errorf("Expected %s to contain [%s]. Got [%s].", c.file, c.data, data)
		}

		teardown(dir)
	}
}

func setup() string {
	dir, _ := ioutil.TempDir("/tmp", "drone_hg_test_")
	os.Mkdir(dir, 0777)
	return dir
}

func teardown(dir string) {
	os.RemoveAll(dir)
}

func readFile(dir, file string) string {
	filename := filepath.Join(dir, file)
	data, _ := ioutil.ReadFile(filename)
	return string(data)
}
