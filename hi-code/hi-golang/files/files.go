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

package files

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Main() {
	fmt.Println("files...")

	// createFile()

	// createDirectory()

	// readFile()

	// appendFile()

	// renameFile()

	//copyFile()
}

func createFile() {
	fileName := "files.md"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	filePath := wd + "/files/" + fileName
	if _, err = os.Stat(filePath); err == nil {
		log.Fatalf("%v file is exist!", fileName)
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Create %v exception!", fileName)
	}
	content := []byte("test content1!\n")
	len, err := file.Write(content)
	defer file.Close()
	if err != nil {
		log.Fatalf("Write %v exception!", fileName)
	}
	log.Printf("Create %v success!,Lenght:%v \n", fileName, len)

}

func createDirectory() {
	directoryName := "test"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	filePath := wd + "/files/" + directoryName
	if _, err = os.Stat(filePath); err == nil {
		log.Fatalf("%v directory is exist!", directoryName)
	}
	err = os.MkdirAll(wd+"/files/"+directoryName, 0755)
	if err != nil {
		log.Fatalf("Create %v exception!", directoryName)
	}
	log.Printf("Create %v success! \n", directoryName)
}

func readFile() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	filePath := wd + "/files/files.md"
	if _, err = os.Stat(filePath); err != nil {
		log.Fatalf("%v file not exists!\n", filePath)
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil || bytes == nil {
		log.Fatalf("ReadFile %v exception!", filePath)
	}
	// by whole
	whole_content := string(bytes)
	fmt.Println(whole_content)

	// by character
	data := bufio.NewScanner(strings.NewReader(whole_content))
	data.Split(bufio.ScanRunes)
	for data.Scan() {
		fmt.Print(data.Text())
	}
	fmt.Println()

}

//append
func appendFile() {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	filePath := wd + "/files/files.md"
	// b, aa := xx()
	// b1, aa := xx()
	// fmt.Println(b, b1, aa)

	message := "appendFile content!"
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open %v exception!\n", filePath)
	}
	defer f.Close()
	fmt.Fprintln(f, message)

}

// test call function use same variables
// func xx() (xx1 string, xx2 int) {
// 	xx1 = "xx1"
// 	xx2 = 10
// 	return
// }

func renameFile() {
	fileName := "files.md"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	oldPath := wd + "/files/" + fileName
	if _, err = os.Stat(oldPath); err != nil {
		log.Fatalf("%v file is not exist!", oldPath)
	}
	newPath := wd + "/files/" + "new_files.md"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func copyFile() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	var targetFile = new(os.File)
	var sourceFile = new(os.File)
	var written int64
	//var sourceFile, targetFile os.File //

	sourceFileName := wd + "/files/files.go"
	if sourceFile, err = os.Open(sourceFileName); err != nil {
		log.Fatalf("%v file not open!", sourceFileName)
	}
	defer sourceFile.Close()

	targetFileName := wd + "/files/files.txt"
	if targetFile, err = os.Create(targetFileName); err != nil {
		log.Fatalf("%v file create fatal!", sourceFileName)
	}
	defer targetFile.Close()

	if written, err = io.Copy(targetFile, sourceFile); err != nil {
		log.Fatalf("%v file create fatal!", sourceFileName)
	}
	log.Printf("written:%d", written)
}
