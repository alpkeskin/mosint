package modules

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Runner(email string, kind string) {
	if kind == "socialscan" {
		cmd := exec.Command("python3", "modules/SocialScan.py", "-e", email)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		go copyOutput(stdout)
		go copyOutput(stderr)
		cmd.Wait()
	} else if kind == "breachdirectory" {
		cmd := exec.Command("python3", "modules/BreachDirectory.py", "-e", email)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
		go copyOutput(stdout)
		go copyOutput(stderr)
		cmd.Wait()
	} else if kind == "hunter" {
		splt := strings.Split(email, "@")
		cmd := exec.Command("python3", "modules/Hunter.py", "-d", splt[1])
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
		go copyOutput(stdout)
		go copyOutput(stderr)
		cmd.Wait()
	} else if kind == "emailrep" {
		cmd := exec.Command("python3", "modules/EmailRep.py", "-e", email)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
		go copyOutput(stdout)
		go copyOutput(stderr)
		cmd.Wait()
	} else {
		color.Red("Error!")
	}

}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
