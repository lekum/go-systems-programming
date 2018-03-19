package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walkFunction(currentPath string, info os.FileInfo, err error) error {
	fileInfo, _ := os.Lstat(currentPath)
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println("Skipping", currentPath)
		return nil
	}
	fileInfo, err = os.Stat(currentPath)
	if err != nil {
		fmt.Println("*", err)
		return err
	}
	mode := fileInfo.Mode()
	if mode.IsDir() {
		tempPath := strings.Replace(currentPath, _path, "", 1)
		pathToCreate := newPath + "/" + filepath.Base(_path) +
			tempPath
		if *minusTest {
			fmt.Println(":", pathToCreate)
			return nil
		}
		_, err := os.Stat(pathToCreate)
		if os.IsNotExist(err) {
			os.MkdirAll(pathToCreate, permissions)
		} else {
			fmt.Println("Did not create", pathToCreate, ":", err)
		}
	}
	return nil
}

var minusTest = flag.Bool("test", false, "Test run")
var _path string
var newPath string
var permissions os.FileMode

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || len(args) == 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	_path = args[0]
	newPath = args[1]

	permissions = os.ModePerm
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		os.MkdirAll(newPath, permissions)
	} else {
		fmt.Println(newPath, "already exists - quitting...")
		os.Exit(1)
	}

	err = filepath.Walk(_path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
