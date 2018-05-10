package main

import (
	"strings"
	"os"
	"fmt"
	"os/exec"
	"bytes"
	"errors"
	"syscall"
	"strconv"
	"bufio"
)

type process struct {
	name     string
	pid      string
	port     string
	portInfo string
	stance   string
}

func notSpace(word string) bool {
	s := strings.Trim(word, " ")
	return len(s) > 0
}

func exitIfError(e error, m string) {
	if e != nil {
		if m == "" {
			fmt.Println("Error: ", e)
		} else {
			fmt.Println(m)
		}
		os.Exit(1)
	}
}

func choose(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func handleInputErrors() {
	if len(os.Args) < 2 {
		exitIfError(errors.New("No arguments received"), "Usage: psoff <port number>")
	}
}

func (p process) print() string {
	return concat("Process ",p.pid,": [ ",p.name," ] is running on port ",p.port)
}



func concat(words ...string) string {
	var b strings.Builder
	for _,word := range words {
		b.WriteString(word)
	}
	return b.String()
}

func main() {
	handleInputErrors()
	portArg := os.Args[1]
	_, err1 := strconv.ParseInt(portArg, 10, 64)
	exitIfError(err1, concat(portArg, " is not a recognised port number"))

	portInfo := strings.Join([]string{"tcp:", portArg}, "")
	listCmd := exec.Command("lsof", "-i", portInfo)

	var out bytes.Buffer
	listCmd.Stdout = &out

	err := listCmd.Run()
	exitIfError(err, concat("No processes were found running on port ", portArg))

	lines := strings.Split(out.String(), "\n")

	for _, line := range lines[1 : len(lines)-1] {
		words := strings.Split(line, " ")
		cw := choose(words, notSpace)
		proc := process{name: cw[0], pid: cw[1], port: portArg, portInfo: cw[8], stance: cw[9]}
		d := proc.print()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(d)
		fmt.Print("Would you like to kill it? [Y/n] ")

		confirm, readErr := reader.ReadString('\n')
		exitIfError(readErr, "I'm sorry, I don't understand.")
		if confirm == "y\n" || confirm == "\n" || confirm == "Y\n" {
			pid, err := strconv.ParseInt(proc.pid, 10, 32)
			exitIfError(err, "Something went wrong.")
			syscall.Kill(int(pid), syscall.SIGKILL)
			fmt.Println("The process has been killed without mercy!")
		} else {
			fmt.Println("The process has been spared for now.")
		}
	}
	os.Exit(0)
}
