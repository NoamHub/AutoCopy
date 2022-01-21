package main

import (
    "io"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"os"
)

var source_dir = "C:\\Projects\\Go\\AutoCopy\\Source"
var dest_dir = "C:\\Projects\\Go\\AutoCopy\\Dest"
 

func main() {
	items, _ := ioutil.ReadDir(source_dir)
    for _, item := range items {
        if item.IsDir() {
            subitems, _ := ioutil.ReadDir(item.Name())
            for _, subitem := range subitems {
                if !subitem.IsDir() {
                    // handle file there
					//io.Copy(destination, source)
                    fmt.Println("Dict")
                }
            }
        } else {
			fmt.Println(item.Name(), item.ModTime())
			source_path := filepath.Join(source_dir, item.Name())
			dest_path := filepath.Join(dest_dir, item.Name())
	        source, err := os.Open(source_path)
			if err != nil {
					return
			}
			defer source.Close()

			destination, err := os.Create(dest_path)
			if err != nil {
					return
			}
            
			io.Copy(destination, source)
        }
    }
}
