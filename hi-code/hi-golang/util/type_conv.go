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

package util

import (
	"fmt"
	"reflect"
	"strconv"
)

func Main() {

	fmt.Println("util...")
	stringConvTo()

	convToString()

	intToInt()

	charToInt()

	intToChar()
}

func stringConvTo() {
	//string to int
	string_one := "10"
	str_two := "T" //It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	int_one, _ := strconv.Atoi(string_one)
	int_two := int_one
	fmt.Printf("int_one value: %v,"+
		"int_one type: %v,"+
		"int_one point: %v,"+
		"int_two point: %v"+
		"\n", int_one, reflect.TypeOf(int_one), &int_one, &int_two) //int_one value: 10,int_one type: int,int_one point: 0xc000016308,int_two point: 0xc000016310

	unit_one, _ := strconv.ParseUint(string_one, 0, 0)
	fmt.Printf("unit_one value: %v,"+
		"unit_one type: %v,"+
		"unit_one point: %v"+
		"\n", unit_one, reflect.TypeOf(unit_one), &unit_one) //unit_one value: 10,unit_one type: uint64,unit_one point: 0xc000016318

	int64_one, _ := strconv.ParseInt(string_one, 0, 0)
	fmt.Printf("int64_one value: %v,"+
		"int64_one type: %v,"+
		"int64_one poin: t%v"+
		"\n", int64_one, reflect.TypeOf(int64_one), &int64_one) //int64_one value: 10,int64_one type: int64,int64_one poin: t0xc000016320

	float64_one, _ := strconv.ParseFloat(string_one, 32)
	fmt.Printf("float64_one value: %v,"+
		"float64_one type: %v,"+
		"float64_one point: %v"+
		"\n", float64_one, reflect.TypeOf(float64_one), &float64_one) //float64_one value: 10,float64_one type: float64,float64_one point: 0xc000016328

	float64_two, _ := strconv.ParseFloat("1.0640000343e+01", 64)
	fmt.Printf("float64_two value: %v,"+
		"float64_two type: %v,"+
		"float64_two point: %v"+
		"\n", float64_two, reflect.TypeOf(float64_two), &float64_two) //float64_two value: 10.640000343,float64_two type: float64,float64_two point: 0xc000016340

	complex128_one, _ := strconv.ParseComplex(string_one, 64)
	fmt.Printf("complex128_one value: %v,"+
		"complex128_one type: %v,"+
		"complex128_one point: %v"+
		"\n", complex128_one, reflect.TypeOf(complex128_one), &complex128_one) //complex128_one value: (10+0i),complex128_one type: complex128,complex128_one point: 0xc000016360

	bool_one, _ := strconv.ParseBool(str_two)
	fmt.Printf("bool_one value: %v,"+
		"bool_one type: %v,"+
		"bool_one point: %v"+
		"\n", bool_one, reflect.TypeOf(bool_one), &bool_one) //bool_one value: true,bool_one type: bool,bool_one point: 0xc000016358

}

func convToString() {
	bool_one := true
	bool_two := false

	string_one := strconv.FormatBool(bool_one)
	string_two := strconv.FormatBool(bool_two)

	fmt.Printf("bool_one value: %v,"+
		"bool_two value: %v,"+
		"string_one value: %v,"+
		"string_two value: %v,"+
		"string_one point: %v"+
		"\n", bool_one, bool_two, string_one, string_two, &string_one) //bool_one value: true,bool_two value: false,string_one value: true,string_two value: false,string_one point: 0xc000096710

	var complex128_one complex128 = (10 + 0i)
	string_two = strconv.FormatComplex(complex128_one, 'E', 10, 128)
	fmt.Printf("complex128 :%v To string_two value: %v"+"\n", complex128_one, string_two) //complex128 :(10+0i) To string_two value: (1.0000000000E+01+0.0000000000E+00i)

	float64_one := 10.640000343
	string_two = strconv.FormatFloat(float64_one, 'e', 10, 64)
	fmt.Printf("float64_one :%v To string_two value: %v"+"\n", float64_one, string_two) //float64_one :10.640000343 To string_two value: 1.0640000343e+01

	var int64_one int64 = 60
	string_two = strconv.FormatInt(int64_one, 10)
	fmt.Printf("int64_one :%v To string_two value: %v"+"\n", int64_one, string_two) //int64_one :60 To string_two value: 60

	int64_one = 60
	string_two = strconv.FormatInt(int64_one, 11)                                   // 2 <= base <= 36
	fmt.Printf("int64_one :%v To string_two value: %v"+"\n", int64_one, string_two) //int64_one :60 To string_two value: 55

	var uint64_one uint64 = 60
	string_two = strconv.FormatUint(uint64_one, 10)
	fmt.Printf("uint64_one :%v To string_two value: %v"+"\n", uint64_one, string_two) //uint64_one :60 To string_two value: 60

	uint64_one = 60
	string_two = strconv.FormatUint(uint64_one, 11)
	fmt.Printf("uint64_one :%v To string_two value: %v"+"\n", uint64_one, string_two) //uint64_one :60 To string_two value: 55

	uint64_one = 60
	string_two = strconv.FormatUint(uint64_one, 12)
	fmt.Printf("uint64_one :%v To string_two value: %v"+"\n", uint64_one, string_two) //uint64_one :60 To string_two value: 50

	uint64_one = 60
	string_two = strconv.FormatUint(uint64_one, 13)
	fmt.Printf("uint64_one :%v To string_two value: %v"+"\n", uint64_one, string_two) //uint64_one :60 To string_two value: 48

	uint64_one = 60
	string_two = strconv.FormatUint(uint64_one, 36)
	fmt.Printf("uint64_one :%v To string_two value: %v"+"\n", uint64_one, string_two) //uint64_one :60 To string_two value: 1o

}

func intToInt() {

	var int64_one int = -32769
	int16_one := int16(int64_one)                                                       //Range: -32768 through 32767.
	fmt.Printf("int64_one value :%v To int16_one value: %v"+"\n", int64_one, int16_one) //int64_one value :-32769 To int16_one value: 32767

	int64_one = 32769
	int16_one = int16(int64_one)
	fmt.Printf("int64_one value :%v To int16_one value: %v"+"\n", int64_one, int16_one) //int64_one value :32769 To int16_one value: -32767

	int64_one = -32769
	int32_one := int32(int64_one)
	fmt.Printf("int64_one value :%v To int32_one value: %v"+"\n", int64_one, int32_one) //int64_one value :-32769 To int32_one value: -32769

	int64_one = -32769
	int64_two := int64(int64_one)
	fmt.Printf("int64_one value :%v To int64_two value: %v"+"\n", int64_one, int64_two) //int64_one value :-32769 To int64_one value: -32769

	int64_one = 2147483647
	float32_one := float32(int64_one)
	fmt.Printf("int64_one value :%v To float32_one value: %v"+"\n", int64_one, float32_one) //int64_one value :2147483647 To float32_one value: 2.1474836e+09

	int64_one = -9223372036854775807
	float64_one := float64(int64_one)
	fmt.Printf("int64_one value :%v To float64_one value: %v"+"\n", int64_one, float64_one) //int64_one value :-9223372036854775807 To float64_one value: -9.223372036854776e+18

}

func charToInt() {
	fmt.Println("charToInt...")
	rune_one := 'A'
	fmt.Println("rune_one before", rune_one, "rune_one after", string(rune_one))
}

func intToChar() {
	fmt.Println("intToChar...")
	int_one := 65
	fmt.Println(rune(int_one))
}
