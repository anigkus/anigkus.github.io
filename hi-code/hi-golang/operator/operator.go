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

import (
	"fmt"
	"reflect"
)

func Main() {
	fmt.Println("operator...")

	arithmetic()

	assignment()

	comparison()

	logical()

	bitwise()
}

func arithmetic() {
	fmt.Println("arithmetic")

	int64_one := 20
	int64_one = int64_one + 10
	fmt.Println(int64_one) //30
	int64_one = int64_one - 10
	fmt.Println(int64_one) //20
	int64_one = int64_one * 10
	fmt.Println(int64_one) //200
	int64_one = int64_one / 10
	fmt.Println(int64_one) //20
	int64_one = int64_one % 10
	fmt.Println(int64_one) //0
	int64_one++
	fmt.Println(int64_one) //1
	int64_one--
	fmt.Println(int64_one) //0

}

func assignment() {
	fmt.Println("assignment")
	var int64_one int64 = 20
	fmt.Println(int64_one) //20
	int64_two := 20
	fmt.Println(int64_two) //20
	var int64_three, int64_four = 20, 30
	int64_five := int64_four
	fmt.Println(int64_three, int64_four, int64_five) //20,30,30

}
func comparison() {
	fmt.Println("comparison")
	int64_one := 30
	int64_two := 20
	result := int64_one > int64_two
	fmt.Println("result", result, "result type", reflect.TypeOf(result)) //result true result type bool
	result = int64_one >= int64_two
	fmt.Println(result) //true
	result = int64_one < int64_two
	fmt.Println(result) //false
	result = int64_one <= int64_two
	fmt.Println(result) //false
	result = int64_one != int64_two
	fmt.Println(result) //true
	result = int64_one == int64_two
	fmt.Println(result) //false

}
func logical() {
	fmt.Println("logical")
	int64_one := 30
	int64_two := 20
	result := int64_one > 10 && int64_two > 10
	fmt.Println("result", result, "result type", reflect.TypeOf(result)) //result true result type bool
	result = int64_one > 10 || int64_two > 10
	fmt.Println(result) //true
	result = !(int64_one > 10 && int64_two > 20)
	fmt.Println(result) //true
	result = !(int64_one > 10 || int64_two > 10)
	fmt.Println(result) //false
}
func bitwise() {
	fmt.Println("bitwise")
	int64_one := 30 //11110
	int64_two := 20 //10100
	//00010100=0*2^7+0*2^6+0*2^5+1*2^4+0*2^3+1*2^2+0*2^1+0*2^0=0+0+0+16+0+4+0+0=20
	result := int64_one & int64_two                                      //Sets each bit to 1 if both bits are 1
	fmt.Println("result", result, "result type", reflect.TypeOf(result)) //result 20 result type bool
	//00011110=0*2^7+0*2^6+0*2^5+1*2^4+1*2^3+1*2^2+1*2^1+0*2^0=0+0+0+16+8+4+2+0=30
	result = int64_one | int64_two //Sets each bit to 1 if one of two bits is 1
	fmt.Println(result)            //30
	//00001010=0*2^7+0*2^6+0*2^5+0*2^4+1*2^3+0*2^2+1*2^1+0*2^0=0+0+0+0+8+0+2+0=10
	result = int64_one ^ int64_two //Sets each bit to 1 if only one of two bits is 1
	fmt.Println(result)            //10
	//01111000=0*2^7+1*2^6+1*2^5+1*2^4+1*2^3+0*2^2+0*2^1+0*2^0=0+64+32+16+8+0+0+0=120
	result = int64_one << 2 //Shift left by pushing zeros in from the right and let the leftmost bits fall off
	fmt.Println(result)     //120
	//00000111=0*2^7+0*2^6+0*2^5+0*2^4+0*2^3+1*2^2+1*2^1+1*2^0=0+0+0+0+0+4+2+1=7
	result = int64_one >> 2 //Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off
	fmt.Println(result)     //7
}
