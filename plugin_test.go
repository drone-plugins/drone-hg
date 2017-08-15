package main

import (
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
		event:  "push",
		branch: "default",
		commit: "37526193d0139f188b20e5c8bed8fc0640c38627",
		file:   "test",
		data:   "",
	},
	// head commit
	{
		path:   "cedk/drone-hg-test",
		clone:  "https://bitbucket.org/cedk/drone-hg-test",
		event:  "push",
		branch: "default",
		commit: "6e4637b5ef305536115bd05e97c913fc7cdcc69d",
		file:   "test",
		data:   "Test\n",
	},
}

// TestClone tests the ability to clone a specific commit into
// a fresh, empty directory every time.
func TestClone(t *testing.T) {
	for _, c := range commits {
		dir := setup()
		defer teardown(dir)

		plugin := Plugin{
			Repo: Repo{
				Clone: c.clone,
			},
			Build: Build{
				Path:   filepath.Join(dir, c.path),
				Commit: c.commit,
				Event:  c.event,
			},
		}

		if err := plugin.Exec(); err != nil {
			t.Errorf("Expected successful clone. Got error. %s.", err)
		}

		data := readFile(plugin.Build.Path, c.file)
		if data != c.data {
			t.Errorf("Expected %s to contain [%s]. Got [%s].", c.file, c.data, data)
		}
	}
}

// TestCloneNonEmpty tests the ability to clone a specific commit into
// a non-empty directory. This is useful if the mercurial workspace is cached
// and re-stored for every build.
func TestCloneNonEmpty(t *testing.T) {
	dir := setup()
	defer teardown(dir)

	for _, c := range commits {

		plugin := Plugin{
			Repo: Repo{
				Clone: c.clone,
			},
			Build: Build{
				Path:   filepath.Join(dir, c.path),
				Commit: c.commit,
				Event:  c.event,
			},
		}

		if err := plugin.Exec(); err != nil {
			t.Errorf("Expected successful clone. Got error. %s.", err)
		}

		data := readFile(plugin.Build.Path, c.file)
		if data != c.data {
			t.Errorf("Expected %s to contain [%s]. Got [%s].", c.file, c.data, data)
		}
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
