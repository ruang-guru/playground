package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//TODO: Make sure benchcmp `go install golang.org/x/tools/cmd/benchcmp@latest` or `go install golang.org/x/perf/cmd/...@latest``
func RunBenchmark(benchmarkDir string, fileName string) error {
	var cmd *exec.Cmd
	commandString := fmt.Sprintf("???", benchmarkDir, fileName) // TODO: replace this
	os := runtime.GOOS
	if os == "windows" {
		cmd = exec.Command("powershell", commandString)
	} else {
		cmd = exec.Command("bash", "-c", commandString)
	}
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(stdout))
	return nil
}

func FileExist(fileName string) string {
	if _, err := os.Stat(fileName); err == nil {
		return fileName
	} else {
		return err.Error()
	}
}

func GetGobin() (string, error) {
	command := "go env GOPATH"
	os := runtime.GOOS
	var cmd *exec.Cmd
	var bin string
	if os == "windows" {
		cmd = exec.Command("powershell", command)
		bin = "\\bin"
	} else {
		cmd = exec.Command("bash", "-c", command)
		bin = "/bin"
	}
	stdout, err := cmd.CombinedOutput()
	gobin := strings.ReplaceAll(string(stdout), "\n", "") + bin
	return gobin, err
}

func RunBenchcmp() error {
	var cmd *exec.Cmd
	os := runtime.GOOS
	fileName := "report.txt"
	GoBin, err := GetGobin()
	if err != nil {
		return err
	}
	commandString := fmt.Sprintf("???", GoBin,fileName) // TODO: replace this
	if os == "windows" {
		cmd = exec.Command("powershell", commandString)
	} else {
		cmd = exec.Command("bash", "-c", commandString)
	}
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
