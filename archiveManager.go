package main

import (
	ziplib "archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func zip(source, destination string) {
	files, err := listFiles(source)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		fmt.Println("No files to zip.")
		os.Exit(1)
	}
	archive, err := os.Create(destination)
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := ziplib.NewWriter(archive)
	defer zipWriter.Close()

	for _, file := range files {
		fileToZip, _ := os.Open(file)
		defer fileToZip.Close()

		zipFilePath, _ := filepath.Rel(source, file)

		writer, _ := zipWriter.Create(zipFilePath)
		_, copyErr := io.Copy(writer, fileToZip)
		if copyErr != nil {
			panic(copyErr)
		}
	}
}

func unzip(source, destination string) {
	reader, err := ziplib.OpenReader(source)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	for _, file := range reader.File {
		fileToWrite, _ := os.Create(destination + file.Name)
		defer fileToWrite.Close()

		fileToRead, _ := file.Open()
		defer fileToRead.Close()

		_, copyErr := io.Copy(fileToWrite, fileToRead)
		if copyErr != nil {
			panic(copyErr)
		}
	}
}
