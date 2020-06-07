package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {


	//abs=>absolute path, .=>current directory
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path", root)

	// refer golang filepath doc for walk
	err := filepath.Walk(root, processpath)
	if err != nil{
		fmt.Println("Error: ", err)
	}
}

// refer the func signature of 'WalkFunc' (processpath has same func sig)
func processpath(path string, info os.FileInfo, err error) error{
	if err != nil{
		return err
	}

	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory: ", path)
		}else{
			fmt.Println("File: ", path)
		}
	}
	return nil
}
