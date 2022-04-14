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

package operator

import "fmt"

func Main() {
	fmt.Println("operator...")
	/*
		Arithmetic Operators
		Assignment Operators
		Comparison Operators
		Logical Operators
		Bitwise Operators
	*/
}

func Arithmetic() {

	int64_one := 64
	fmt.Println("Arithmetic "+
		"+", int64_one,
		"-", int64_one,
		"*", int64_one,
		"\\", int64_one,
		"%", int64_one,
		"++", int64_one,
		"--", int64_one)
}

func Assignment() {
	fmt.Println("Assignment")
}
func Comparison() {
	fmt.Println("Comparison")
}
func Logical() {
	fmt.Println("Logical")
}
func Bitwise() {
	fmt.Println("Bitwise")
}
