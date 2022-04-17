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
package defers

import (
	"fmt"
)

func Main() {
	//deferred calls are executed in last-in-first-out order.(LIFO)
	fmt.Println("Defers...")
	defer func() {
		recover := recover()
		fmt.Println("Main Recovered. Error:\n", recover)
	}()

	processPanic()

	afterFunc() //This function will only be called if processPanic is defined in defer function(not NormalFunc defined defer)
}

func afterFunc() {
	fmt.Println("AfterFunc")
}

func processPanic() {
	normalFunc()
	//fmt.Println(returnFunc())
	defer func() {
		fmt.Println("processPanic-defer-func-1-before")
		recover := recover()
		fmt.Println(recover)
		fmt.Println("processPanic-defer-func-1-after")
	}()

	// defer func() {
	// 	fmt.Println("processPanic-defer-func-2-before")
	// 	recover := recover()
	// 	fmt.Println(recover)
	// 	fmt.Println("processPanic-defer-func-2-after")
	// }()

	panic("Process panic problem")
}

func normalFunc() {
	func() {
		xx := recover()
		fmt.Println(xx)
		fmt.Println("NormalFunc-1")
	}()
	func() {
		xx := recover()
		fmt.Println(xx)
		fmt.Println("NormalFunc-2")
	}()
}

func returnFunc() int {
	return func() int {
		fmt.Println("ReturnFunc")
		return 1
	}()
}
