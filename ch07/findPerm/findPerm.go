package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var PERMISSIONS string

func permissionsOfFile(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		return "-1"
	}
	mode := info.Mode()
	return mode.String()[1:10]
}

func walkFunction(path string, info os.FileInfo, err error) error {
	_, err = os.Lstat(path)
	if err != nil {
		return err
	}

	if permissionsOfFile(path) == PERMISSIONS {
		fmt.Println(path)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("usage: %s RootDirectory permissions\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	Path := args[1]
	Path, _ = filepath.EvalSymlinks(Path)
	PERMISSIONS = args[2]

	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
	}

}
