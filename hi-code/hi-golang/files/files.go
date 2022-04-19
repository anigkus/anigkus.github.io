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
	"fmt"
	"log"
	"os"
)

func Main() {
	fmt.Println("files...")

	createFile()

	createDirectory()

	// readFile()

	// writeFile()
}

func createFile() {
	fileName := "files.md"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd %v exception!", fileName)
	}
	filePath := wd + "/files/" + fileName
	if _, err = os.Stat(filePath); err == nil {
		log.Fatalf("%v file is exist!", fileName)
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Create %v exception!", fileName)
	}
	content := []byte("test content!")
	len, err := file.Write(content)
	if err != nil {
		log.Fatalf("Write %v exception!", fileName)
	}
	log.Printf("Create %v success!,Lenght:%v \n", fileName, len)
	defer file.Close()

}

func createDirectory() {
	directoryName := "test"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd %v exception!", directoryName)
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

}

func writeFile() {

}
