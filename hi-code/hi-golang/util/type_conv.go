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
	//string to int
	str_one := "10"
	str_two := "T" //It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	int_one, _ := strconv.Atoi(str_one)
	int_two := int_one
	fmt.Printf("int_one value: %v,"+
		"int_one type: %v,"+
		"int_one point: %v,"+
		"int_two point: %v,"+
		"\n", int_one, reflect.TypeOf(int_one), &int_one, &int_two)

	unit_one, _ := strconv.ParseUint(str_one, 0, 0)
	fmt.Printf("unit_one value: %v,"+
		"unit_one type: %v,"+
		"unit_one point: %v,"+
		"\n", unit_one, reflect.TypeOf(unit_one), &unit_one)

	int64_one, _ := strconv.ParseInt(str_one, 0, 0)
	fmt.Printf("int64_one value: %v,"+
		"int64_one type: %v,"+
		"int64_one poin: t%v,"+
		"\n", int64_one, reflect.TypeOf(int64_one), &int64_one)

	float64_one, _ := strconv.ParseFloat(str_one, 32)
	fmt.Printf("float64_one value: %v,"+
		"float64_one type: %v,"+
		"float64_one point: %v,"+
		"\n", float64_one, reflect.TypeOf(float64_one), &float64_one)

	float64_two, _ := strconv.ParseFloat(str_one, 64)
	fmt.Printf("float64_two value: %v,"+
		"float64_two type: %v,"+
		"float64_two point: %v,"+
		"\n", float64_two, reflect.TypeOf(float64_two), &float64_two)

	complex128_one, _ := strconv.ParseComplex(str_one, 64)
	fmt.Printf("complex128_one value: %v,"+
		"complex128_one type: %v,"+
		"complex128_one point: %v,"+
		"\n", complex128_one, reflect.TypeOf(complex128_one), &complex128_one)

	bool_one, _ := strconv.ParseBool(str_two)
	fmt.Printf("bool_one value: %v,"+
		"bool_one type: %v,"+
		"bool_one point: %v,"+
		"\n", bool_one, reflect.TypeOf(bool_one), &bool_one)

}
