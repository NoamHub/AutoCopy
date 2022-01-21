package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var source_dir = "C:\\Projects\\Go\\AutoCopy\\Source"
var dest_dir = "C:\\Projects\\Go\\AutoCopy\\Dest"
var timestamp_path = "C:\\Projects\\Go\\AutoCopy\\TS"

var file_types = map[string]bool{"jpg": true}
var files_in_dest_dir = map[string]bool{}

func main() {

	file, err := os.Stat(timestamp_path)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Timestamp file is missing. Aborting")
		return
	}

	timestamp := file.ModTime()
	//fmt.Println(ts.Before(time.Now()))

	ScanFolder(source_dir, timestamp)
}

func ScanFolder(directory string, timestamp time.Time) {
	fmt.Println("Scanning:", directory)

	items, _ := ioutil.ReadDir(directory)
	for _, item := range items {
		if item.IsDir() {
			fmt.Println("Sub-directory detected:", item.Name())
			ScanFolder(filepath.Join(directory, item.Name()), timestamp)
		} else {
			fmt.Println(item.Name(), item.ModTime())
			if files_in_dest_dir[item.Name()] {
				panic("File already found in destination directory")
			}

			source_path := filepath.Join(directory, item.Name())
			dest_path := filepath.Join(dest_dir, item.Name())
			source, err := os.Open(source_path)
			if err != nil {
				panic(err)
			}
			defer source.Close()

			destination, err := os.Create(dest_path)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(destination, source)
			if err != nil {
				panic(err)
			}
	

			files_in_dest_dir[item.Name()] = true
		}
	}
}

func ScanDestFolder(directory string, found_files map[string]bool) {
	fmt.Println("Scanning destination directory")
	files, err := ioutil.ReadDir(dest_dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		found_files[file.Name()] = true
	}
}
