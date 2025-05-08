package main

import (
	"anvil/cli"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func main() {
	// get home dir
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// create anvil-cli dir in homedir
	newpath := filepath.Join(dirname, ".anvil-cli")
	// err = os.MkdirAll(newpath, os.ModePerm)
	if err := ensureDir(newpath); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	fmt.Println(dirname)
	os.Exit(cli.CLI())
}
