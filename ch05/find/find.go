package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var minusS = flag.Bool("s", false, "Sockets")
var minusP = flag.Bool("p", false, "Pipes")
var minusSL = flag.Bool("sl", false, "Symbolic links")
var minusD = flag.Bool("d", false, "Directories")
var minusF = flag.Bool("f", false, "Files")

var printAll bool

func walkFunction(path string, info os.FileInfo, err error) error {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if printAll {
		fmt.Println(path)
		return nil
	}
	mode := fileInfo.Mode()
	if mode.IsRegular() && *minusF {
		fmt.Println(path)
		return nil
	}
	if mode.IsDir() && *minusD {
		fmt.Println(path)
		return nil
	}
	fileInfo, _ = os.Lstat(path)
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		if *minusSL {
			fmt.Println(path)
			return nil
		}
	}
	if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		if *minusP {
			fmt.Println(path)
			return nil
		}
	}
	if fileInfo.Mode()&os.ModeSocket != 0 {
		if *minusS {
			fmt.Println(path)
			return nil
		}
	}
	return nil
}

func main() {

	flag.Parse()
	args := flag.Args()

	printAll = false

	if *minusS && *minusP && *minusSL && *minusD && *minusF {
		printAll = true
	}

	if !(*minusS || *minusP || *minusSL || *minusD || *minusF) {
		printAll = true
	}

	if len(args) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	Path := args[0]
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
