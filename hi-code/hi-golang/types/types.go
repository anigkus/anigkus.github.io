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

package types

import (
	"fmt"
)

//Episode 1
func Main() {

	variable_declaration_style()

	variable_declaration_bool()

	variable_declaration_number()

}

func variable_declaration_style() {

	fmt.Printf("variable_declaration_style\n")

	var v_int_one int = 10
	var v_int_two = 20
	v_int_three := 30 //Infer the type automaticlly

	fmt.Printf("var $name type = $value\t%v \n", v_int_one)
	fmt.Printf("var $name = $value\t%v \n", v_int_two)
	fmt.Printf("var $name:= $value\t%v \n", v_int_three)

	fmt.Printf("\n")
}

func variable_declaration_bool() {

	fmt.Printf("variable_declaration_bool\n")
	var v_bool_one bool = true
	var v_bool_two bool = true
	var v_bool_three bool = v_bool_one
	var v_bool_four bool = false
	var v_bool_five bool = false
	var v_bool_six bool = v_bool_four
	var v_bool_seven bool = false

	fmt.Println("&v_bool_one:", &v_bool_one, "&v_bool_two:", &v_bool_two, "&v_bool_three:", &v_bool_three)
	fmt.Println("&v_bool_four:", &v_bool_four, "&v_bool_five:", &v_bool_five, "&v_bool_six:", &v_bool_six, "&v_bool_seven:", &v_bool_seven)
	fmt.Println("v_bool_one==v_bool_two:", v_bool_one == v_bool_two)
	fmt.Println("v_bool_three==v_bool_one:", v_bool_three == v_bool_one)
	fmt.Println("v_bool_four==v_bool_five:", v_bool_four == v_bool_five)
	fmt.Println("v_bool_six==v_bool_four:", v_bool_six == v_bool_four)
	fmt.Println("v_bool_three===v_bool_one:", &v_bool_three == &v_bool_one)
	fmt.Println("v_bool_six===v_bool_four:", &v_bool_six == &v_bool_four)

	fmt.Printf("\n")
}

func variable_declaration_number() {

	fmt.Printf("variable_declaration_number\n")
	//unsigned
	var v_uint_min uint = 0
	var v_uint_max uint = 18446744073709551615 //math.MaxUint
	var v_uint8_min uint8 = 0
	var v_uint8_max uint8 = 255
	var v_byte_min byte = v_uint8_min
	var v_byte_max byte = v_uint8_max
	var v_uint16_min uint16 = 0
	var v_uint16_max uint16 = 65535
	var v_uint32_min uint32 = 0
	var v_uint32_max uint32 = 4294967295
	var v_uint64_min uint64 = 0
	var v_uint64_max uint64 = 18446744073709551615

	fmt.Println("&v_uint_min:", &v_uint_min, "&v_uint_max:", &v_uint_max)
	fmt.Println("v_uint_min uint =", v_uint_min, "v_uint_max uint = ", v_uint_max)

	fmt.Println("&v_uint8_min:", &v_uint8_min, "&v_uint8_max:", &v_uint8_max)
	fmt.Println("v_uint8_min uint8 =", v_uint8_min, "v_uint8_max uint8 = ", v_uint8_max)

	fmt.Println("&v_byte_min:", &v_byte_min, "&v_byte_max:", &v_byte_max)
	fmt.Println("v_byte_min byte =", v_byte_min, "v_byte_max byte = ", v_byte_max)

	fmt.Println("&v_uint16_min:", &v_uint16_min, "&v_uint16_max:", &v_uint16_max)
	fmt.Println("v_uint16_min uint16 =", v_uint16_min, "v_uint16_max uint16 = ", v_uint16_max)

	fmt.Println("&v_uint32_min:", &v_uint32_min, "&v_uint32_max:", &v_uint32_max)
	fmt.Println("v_uint32_min uint32 =", v_uint16_min, "v_uint32_max uint32 = ", v_uint32_max)

	fmt.Println("&v_uint64_min:", &v_uint64_min, "&v_uint64_max:", &v_uint64_max)
	fmt.Println("v_uint64_min uint64 =", v_uint64_min, "v_uint64_max uint64 = ", v_uint64_max)

	//signed
	var v_int_min int = 0
	var v_int_max int = 9223372036854775807 //math.MaxInt
	var v_int8_min int8 = -128
	var v_int8_max int8 = 127
	var v_int16_min int16 = -32768
	var v_int16_max int16 = 32767
	var v_int32_min int32 = -2147483648
	var v_int32_max int32 = 2147483647
	var v_rune_min rune = v_int32_min
	var v_rune_max rune = v_int32_max
	var v_int64_min int64 = -9223372036854775808
	var v_int64_max int64 = 9223372036854775807

	var v_float32_min float32 = 1e-46                   //math.SmallestNonzeroFloat32
	var v_float32_max float32 = 3.4028235e+37           //math.MaxFloat32
	var v_float64_min float64 = 5e-324                  //math.SmallestNonzeroFloat64
	var v_float64_max float64 = 1.7976931348623157e+308 //math.MaxFloat64
	var v_complex64_min complex64 = 1e-46 + 1e-46
	var v_complex64_max complex64 = 3.4028235e+37 + 3.4028235e+37 + 3.4028235e+37
	// var v_complex128_min complex128 = -128
	// var v_complex128_max complex128 = 127

	fmt.Println("&v_int_min:", &v_int_min, "&v_int_max:", &v_int_max)
	fmt.Println("v_int_min int =", v_int_min, "v_int_max int = ", v_int_max)

	fmt.Println("&v_int8_min:", &v_int8_min, "&v_int8_max:", &v_int8_max)
	fmt.Println("v_int8_min int8 =", v_int8_min, "v_int8_max int8 = ", v_int8_max)

	fmt.Println("&v_int16_min:", &v_int16_min, "&v_int16_max:", &v_int16_max)
	fmt.Println("v_int16_min int16 =", v_int16_min, "v_int16_max int16 = ", v_int16_max)

	fmt.Println("&v_int32_min:", &v_int32_min, "&v_int32_max:", &v_int32_max)
	fmt.Println("v_int32_min int32 =", v_int32_min, "v_int32_max int32 = ", v_int32_max)

	fmt.Println("&v_rune_min:", &v_rune_min, "&v_rune_max:", &v_rune_max)
	fmt.Println("v_rune_min rune =", v_rune_min, "v_rune_max rune = ", v_rune_max)

	fmt.Println("&v_int64_min:", &v_int64_min, "&v_int64_max:", &v_int64_max)
	fmt.Println("v_int64_min int64 =", v_int64_min, "v_int64_max int64 = ", v_int64_max)

	fmt.Println("&v_float32_min:", &v_float32_min, "&v_float32_max:", &v_float32_max)
	fmt.Println("v_float32_min float32 =", v_float32_min, "v_float32_max float32 = ", v_float32_max)

	fmt.Println("&v_float64_min:", &v_float64_min, "&v_float64_max:", &v_float64_max)
	fmt.Println("v_float64_min float64 =", v_float64_min, "v_float64_max float64 = ", v_float64_max)

	fmt.Println("&v_complex64_min:", &v_complex64_min, "&v_complex64_max:", &v_complex64_max)
	fmt.Println("v_complex64_min complex64 =", v_complex64_min, "v_complex64_max complex64 = ", v_complex64_max)

	fmt.Printf("\n")
}
