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

package loop

import (
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("loop...")
	loopStatement()
}

func loopStatement() {
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], ",") //1,2,3,4,5,
	}
	fmt.Println("\n", strings.Repeat("-", 20))

	for i := range array {
		fmt.Print(array[i], ",") //1,2,3,4,5,
	}
	fmt.Println("\n", strings.Repeat("-", 20))

	slice_one := []string{"10", "", "30", "40", "50"}
	slice_one = append(slice_one, "60", "70")

	// While loop same
	i := 0
	for i < 7 {
		if slice_one[i] != "" {
			fmt.Print(slice_one[i], ",") //10,30,40,50,60,70,
		}
		i++
	}
	fmt.Println("i:", i) //Why
	for i == 7 {
		fmt.Println("00")
		break
	}

	fmt.Println("\n", strings.Repeat("-", 20))
	var map_one = map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	map_one["key4"] = "value4"
	for key, value := range map_one {
		fmt.Println("Key:", key, "=>", "value:", value)
		/*
			Key: key1 => value: value1
			Key: key2 => value: value2
			Key: key3 => value: value3
			Key: key4 => value: value4
		*/
	}
	fmt.Println("\n", strings.Repeat("-", 20))

}
