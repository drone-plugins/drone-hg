package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "mercurial plugin"
	app.Usage = "mercurial plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "remote",
			Usage:  "hg remote url",
			EnvVar: "DRONE_REMOTE_URL",
		},
		cli.StringFlag{
			Name:   "path",
			Usage:  "hg clone path",
			EnvVar: "DRONE_WORKSPACE",
		},
		cli.StringFlag{
			Name:   "sha",
			Usage:  "hg commit rev",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.StringFlag{
			Name:   "netrc.machine",
			Usage:  "netrc machine",
			EnvVar: "DRONE_NETRC_MACHINE",
		},
		cli.StringFlag{
			Name:   "netrc.username",
			Usage:  "netrc username",
			EnvVar: "DRONE_NETRC_USERNAME",
		},
		cli.StringFlag{
			Name:   "netrc.password",
			Usage:  "netrc password",
			EnvVar: "DRONE_NETRC_PASSWORD",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Clone: c.String("remote"),
		},
		Build: Build{
			Path:   c.String("path"),
			Event:  c.String("event"),
			Commit: c.String("sha"),
		},
		Netrc: Netrc{
			Machine:  c.String("netrc.machine"),
			Login:    c.String("netrc.username"),
			Password: c.String("netrc.password"),
		},
	}

	return plugin.Exec()
}
