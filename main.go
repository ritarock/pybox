package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	var file string
	var version float64 = 3.7

	app := &cli.App{
		Name:  "pybox",
		Usage: "Run pthon",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Usage:       "target file",
				Required:    true,
				Aliases:     []string{"f"},
				Destination: &file,
			},
			&cli.Float64Flag{
				Name:        "version",
				Usage:       "target version",
				DefaultText: "3.7",
				Aliases:     []string{"v"},
				Destination: &version,
			},
		},
		Action: func(c *cli.Context) error {
			run(file, version)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(file string, version float64) {
	if version == 0 {
		version = 3.7
	}
	command := createCommand(file, version)

	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}

func createCommand(file string, version float64) string {
	v := map[float64]string{
		3.5: "py35/",
		3.6: "py36/",
		3.7: "py37/",
		3.8: "py38/",
		3.9: "py39/",
	}

	return "FILE=" + file + " docker-compose -f " + v[version] + "docker-compose.yml up"
}
