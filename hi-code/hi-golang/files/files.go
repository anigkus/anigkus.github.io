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
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	//console
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.Println("init started")
}

func Main() {
	fmt.Println("files...")

	//createFile()

	// createDirectory()

	// readFile()

	// appendFile()

	// renameFile()

	//copyFile()

	//fileMetadata()

	//deleteFile()

	//truncateFile()

	//changeFile()
}

func createFile() {
	fileName := "files.md"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	filePath := wd + "/files/" + fileName
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

func fileMetadata() {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	var sourceFile fs.FileInfo
	sourceFileName := wd + "/files/files.go"
	if sourceFile, err = os.Stat(sourceFileName); err != nil {
		log.Fatalf("%v file not exists!", sourceFileName)
	}
	fmt.Println(sourceFile.IsDir())
	fmt.Println(sourceFile.ModTime())
	fmt.Println(sourceFile.Mode())
	fmt.Println(sourceFile.Name())
	fmt.Println(sourceFile.Size())
	fmt.Println(sourceFile.Sys())
	//time := sourceFile.ModTime().Add(time.Duration(time.Duration.Seconds(10)))
	time := sourceFile.ModTime().Add(10 * time.Second)
	fmt.Println("time:", time) // compare sourceFile.ModTime()+10 second
}

func deleteFile() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	sourceFileName := wd + "/files/files.md"
	if err = os.Remove(sourceFileName); err != nil {
		log.Fatalf("%v file remove exception!", sourceFileName)
	} else {
		log.Printf("%v file remove success!\n", sourceFileName)
	}
}

func truncateFile() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	sourceFileName := wd + "/files/files.md"
	if err = os.Truncate(sourceFileName, 3); err != nil {
		log.Fatalf("%v file truncate exception!", sourceFileName)
	} else {
		log.Printf("%v file truncate success!\n", sourceFileName)
	}
}

func changeFile() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd exception!")
	}
	var fileInfo os.FileInfo
	sourceFileName := wd + "/files/files.md"
	if fileInfo, err = os.Stat(sourceFileName); err != nil {
		log.Fatalf("%v file not exists!", sourceFileName)
	}
	//Chown changes the numeric uid and gid of the named file.
	if err = os.Chown(sourceFileName, os.Getuid(), os.Getgid()); err != nil {
		log.Fatalf("%v file change Chown  exception!", err)
	}
	//Chmod changes the mode of the named file to mode.
	if err = os.Chmod(sourceFileName, 0777); err != nil {
		log.Fatalf("%v file change Chmod 0777 exception!", err)
	}
	var atime, mtime time.Time
	fmt.Printf("%v old modTime: %v\n", sourceFileName, fileInfo.ModTime())

	atime = time.Now().Add(time.Hour * 24)
	mtime = time.Now().Add(time.Hour * 48) //When atime!=mtime,by mtime finally
	fmt.Printf("atime:%v,mtime: %v\n", atime, mtime)
	if err = os.Chtimes(sourceFileName, atime, mtime); err != nil {
		log.Fatalf("%v file change atime and  mtime exception!", err)
	}
	if fileInfo, err = os.Stat(sourceFileName); err != nil {
		log.Fatalf("%v file not exists!", err)
	}
	fmt.Printf("%v new modTime: %v\n", sourceFileName, fileInfo.ModTime())
}
