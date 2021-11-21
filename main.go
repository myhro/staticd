package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/myhro/staticd/tools"
)

func exit(msg string) {
	io.WriteString(os.Stderr, fmt.Sprintf("Error: %v\n", msg))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		exit(fmt.Sprintf("usage: %v <tool>", path.Base(os.Args[0])))
	}

	tool := &tools.Tool{
		Name: os.Args[1],
	}

	err := tool.SetURL()
	if err != nil {
		exit(err.Error())
	}

	err = tool.SetRuntime(runtime.GOARCH, runtime.GOOS)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println("Checking version for", tool.Name)
	err = tool.GetVersion()
	if err != nil {
		exit(err.Error())
	}

	err = tool.SetArchive()
	if err != nil {
		exit(err.Error())
	}

	fmt.Println("Downloading version", tool.Version)
	filename, err := tool.Download()
	if err != nil {
		exit(err.Error())
	}
	defer os.Remove(filename)
	fmt.Println("Downloaded", filename)
}
