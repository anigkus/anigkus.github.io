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

package funcs

import (
	"fmt"
)

func Main() {
	fmt.Println("function...")

	internalMethodNoArgsNoReturn()

	fmt.Println(internalMethodNoArgsWithReturn())

	string_one, int64_one, _ := internalMethodNoArgsWithReturns()
	fmt.Println("internalMethodNoArgsWithReturns ", "string_one:", string_one, "int64_one:", int64_one)

	string_two, int64_two, float64_two := internalMethodWithArgsWithReturns("string_one", 100, 200)
	fmt.Println("internalMethodWithArgsWithReturns ", "string_two:", string_two, "int64_two:", int64_two, "float64_two:", float64_two)
}

func internalMethodNoArgsNoReturn() {
	fmt.Println("internalMethodNoArgsNoReturn")
}

func internalMethodNoArgsWithReturn() string {
	return "internalMethodNoArgsWithReturn"
}

func internalMethodNoArgsWithReturns() (string_one string, int64_one int64, float64_one float64) {
	string_one = "string_one"
	int64_one = 100
	float64_one = 200
	return
}

func internalMethodWithArgsWithReturns(string_one string, int64_one int64, float64_one float64) (string_two string, int64_two int64, float64_two float64) {
	string_two = string_one + "|string_two"
	int64_two = int64_one + 100
	float64_two = float64_one + 200
	return
}

func PublicMethodNoArgsNoReturn() {
	fmt.Println("publicMethodNoArgsNoReturn")
}
