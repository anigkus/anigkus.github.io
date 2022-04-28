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
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Main() {
	fmt.Println("file operator...")

	//compressZip()

	//readZip() //Nomal

	//extractZip()

	parseText()
	writeText()

	parseCSV()
	writeCSV()

	parseXML()
	writeXML()

	// parseJSON()
	// writeJSON()
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

func getFileString(methodName string, fileName string) string {
	fmt.Println(methodName)
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Excute %v Getwd Err:%s", methodName, err)
	}
	filePath := pwd + "/fileo/" + fileName
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("Excute %v Stat Err:%s", methodName, err)
	}
	if err != nil {
		log.Printf("Excute %v Stat Err:%s", methodName, err)
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil || bytes == nil {
		log.Printf("Excute %v ReadFile Err:%s", methodName, err)
	}
	return string(bytes) //byte to string
}

func IsNotExist(methodName string, fileName string) bool {
	fmt.Println(methodName)
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Excute %v Getwd Err:%s", methodName, err)
		return false
	}
	filePath := pwd + "/fileo/" + fileName
	if _, err = os.Stat(filePath); errors.Is(err, fs.ErrNotExist) {
		return true
	}
	log.Printf("Excute %v Stat Err:%s", methodName, err)
	return false
}

func parseText() {
	methodName := "parseText"
	stringText := getFileString(methodName, "fileo.txt")

	// by character
	scanner := bufio.NewScanner(strings.NewReader(stringText))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
	fmt.Println("\n" + strings.Repeat("-", 20))
}

func writeText() {
	methodName := "writeText"
	if !IsNotExist(methodName, "fileo_write.txt") {
		return
	}
	pwd, _ := os.Getwd()
	filePath := pwd + "/fileo/fileo_write.txt"
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Excute %v Create Err:%s", methodName, err)
	}
	defer file.Close()
	stringText := "1001,\"Angel\",100\n" +
		"1002,\"Bob\",90\n" +
		"1003,\"Sam\",80"
	n, err := file.WriteString(stringText)
	if err != nil || n <= 0 {
		log.Printf("Excute %v WriteString Err:%s", methodName, err)
	}
	fmt.Println("\n" + strings.Repeat("-", 20))
}

func parseCSV() {
	methodName := "parseCSV"
	// fmt.Println(methodName)
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	log.Printf("Excute %v Getwd Err:%s", methodName, err)
	// }
	// filePath := pwd + "/fileo/fileo.csv"
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Printf("Excute %v Open Err:%s", methodName, err)
	// }
	// defer file.Close()
	// // by Line
	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)
	// for scanner.Scan() {
	// 	fmt.Print(scanner.Text() + "\n")
	// }
	stringText := getFileString(methodName, "fileo.csv")

	// by character
	scanner := bufio.NewScanner(strings.NewReader(stringText))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	fmt.Println("\n" + strings.Repeat("-", 20))

}

func writeCSV() {
	methodName := "writeCSV"
	if !IsNotExist(methodName, "fileo_write.csv") {
		return
	}
	pwd, _ := os.Getwd()
	filePath := pwd + "/fileo/fileo_write.csv"
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Excute %v Create Err:%s", methodName, err)
	}
	stringText := "No,Name,Score\n" +
		"1001,\"Angel\",100\n" +
		"1002,\"Bob\",90\n" +
		"1003,\"Sam\",80"
	n, err := file.WriteString(stringText)
	defer file.Close()
	if err != nil || n <= 0 {
		log.Printf("Excute %v WriteString Err:%s", methodName, err)
	}
	fmt.Println("\n" + strings.Repeat("-", 20))
}

type Student struct {
	No    string `xml:"No"`
	Name  string `xml:"Name"`
	Score string `xml:"Score"`
}

type Students struct {
	Student []Student `xml:"Student"`
}

func parseXML() {
	fileName := "fileo.xml"
	methodName := "parseXML"
	fmt.Println(methodName)
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Excute %v Getwd Err:%s", methodName, err)
	}
	filePath := pwd + "/fileo/" + fileName
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("Excute %v Stat Err:%s", methodName, err)
	}
	if err != nil {
		log.Printf("Excute %v Stat Err:%s", methodName, err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Excute %v ReadFile Err:%s", methodName, err)
	}
	students := &Students{}
	err = xml.Unmarshal([]byte(data), &students)
	if err != nil {
		log.Printf("Excute %v Unmarshal Err:%s", methodName, err)
	}
	if students != nil && students.Student != nil && len(students.Student) > 0 {
		for _, stu := range students.Student {
			fmt.Printf("No:%v,Name:%v,Score:%v\n", stu.No, stu.Name, stu.Score)
		}
	} else {
		log.Printf("Excute %v For Err:%s,students:%v", methodName, err, students)
	}

}

func writeXML() {
	methodName := "writeXML"
	if !IsNotExist(methodName, "fileo_write.xml") {
		return
	}
	pwd, _ := os.Getwd()
	filePath := pwd + "/fileo/fileo_write.xml"
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Excute %v Create Err:%s", methodName, err)
	}
	stringText := "<Students>" +
		"<Student>" +
		"    <No>1001</No>" +
		"   <Name>Angel</Name>" +
		"   <Score>100</Score>" +
		"</Student>" +
		"<Student>" +
		"   <No>1002</No>" +
		"   <Name>Bob</Name>" +
		"   <Score>90</Score>" +
		"</Student>" +
		"<Student>" +
		"   <No>1003</No>" +
		"   <Name>Sam</Name>" +
		"   <Score>80</Score>" +
		"</Student>" +
		"</Students>"
	n, err := file.WriteString(stringText)
	defer file.Close()
	if err != nil || n <= 0 {
		log.Printf("Excute %v WriteString Err:%s", methodName, err)
	}
	fmt.Println("\n" + strings.Repeat("-", 20))
}

func parseJSON() {
	stringText := getFileString("parseJSON", "fileo.json")

	// by character
	data := bufio.NewScanner(strings.NewReader(stringText))
	data.Split(bufio.ScanRunes)
	for data.Scan() {
		fmt.Println(data.Text())
	}
	fmt.Println("\n" + strings.Repeat("-", 20))
}

func writeJSON() {

	if !IsNotExist("writeJSON", "fileo_write.json") {
		return
	}
	//TODO

	fmt.Println("\n" + strings.Repeat("-", 20))
}
