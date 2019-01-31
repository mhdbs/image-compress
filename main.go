package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var flagvar string

	flag.StringVar(&flagvar, "compress", "", "help")
	flag.Parse()
	checkTypeFile(flagvar)
	infile, err := os.Open(flagvar)
	if err != nil {
		log.Printf("failed opening %s:", err)
		panic(err.Error())
	}
	fmt.Println("infile data", infile)
}

func checkTypeFile(file string) {
	fi, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		readFolder(file)
		fmt.Println("directory")
	case mode.IsRegular():
		fmt.Println("file")
	}
}

func readFolder(file string) {
	var files []string
	err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
