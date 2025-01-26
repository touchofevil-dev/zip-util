package main

import (
	"os"
	"path/filepath"
)

func listFiles(source string) ([]string, error) {
	var files []string
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
