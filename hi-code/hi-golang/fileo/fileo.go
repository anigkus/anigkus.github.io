/*
Copyright 2022 The https://github.com/anigkus Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fileo

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Main() {
	fmt.Println("file operator...")

	compressZip()

	//readZip() //Nomal

	//extractZip()
}
func compressZip() {
	fmt.Println("compressZip")
	pwd, _ := os.Getwd()
	filePath := pwd + "/fileo/fileo.zip"

	flags := os.O_WRONLY | os.O_APPEND | os.O_TRUNC
	file, err := os.OpenFile(filePath, flags, 0666)
	if err != nil {
		log.Fatal("OpenFile fileo.zip Exception!")
	}
	defer file.Close()
	//Default relatively execute dirctory
	//var files = []string{pwd + "/fileo/two.md", pwd + "/fileo/three.md"}
	//var files = []string{"README.md", pwd + "/fileo/two.md", pwd + "/fileo/three.md"}
	var files = []string{"main.go"}

	zipw := zip.NewWriter(file)
	defer zipw.Close()
	for _, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}
}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func extractZip() {
	pwd, _ := os.Getwd()
	zipReader, _ := zip.OpenReader(pwd + "/fileo/fileo.zip")
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "/tmp/"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func readZip() {

	pwd, _ := os.Getwd()
	read, err := zip.OpenReader(pwd + "/fileo/fileo.zip")
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for _, file := range read.File {
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}
}

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	fmt.Fprintf(os.Stdout, "%s:", file.Name)

	if err != nil {
		msg := "failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	fmt.Println()

	return nil
}
