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
	"strings"
)

//Episode 1
func Main() {

	variable_declaration_style()

	variable_declaration_bool()

	variable_declaration_number()

	variable_declaration_string()

	variable_declaration_array()

	variable_declaration_slice()

	variable_declaration_map()

	variable_operators_builtin()

	variable_operators_collection()

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

func variable_declaration_string() {

	fmt.Printf("variable_declaration_string\n")
	var v_string_one string = "one\tstring"
	v_string_two := `two\tstring`
	v_string_three := "three" + " string"
	v_string_four := " four "
	var v_string_five string = v_string_four + "five"

	fmt.Println("&v_string_one:", &v_string_one, "&v_string_two:", &v_string_two)
	fmt.Println(
		"v_string_one =", v_string_one,
		"v_string_two =", v_string_two,
		"v_string_three =", v_string_three,
		"v_string_five =", v_string_five)

	//fun
	fmt.Println(len(v_string_one))                          //10
	fmt.Println(strings.Index(v_string_one, "n"))           //1
	fmt.Println(strings.LastIndex(v_string_one, "n"))       //8
	fmt.Println(strings.IndexAny(v_string_one, "n"))        //1
	fmt.Println(strings.LastIndexAny(v_string_one, "n"))    //8
	fmt.Println(strings.Replace(v_string_one, "n", "k", 1)) //oke	string
	fmt.Println(strings.ReplaceAll(v_string_one, "n", "k")) //oke	strikg
	fmt.Println(strings.Split(v_string_three, " "))         //[three string]

	//cutset
	v_cutset_no_trim := "T"
	v_cutset_left_trim := " T"
	v_cutset_right_trim := "T "
	v_cutset_left_and_right_trim := " T "

	v_string_six := "TikHoT"
	v_string_seven := " TikHoT"
	v_string_eight := "TikHoT "
	v_string_nine := " TikHoT "

	fmt.Println(strings.TrimLeft(v_string_nine, " T "))
	//Trim
	fmt.Println(strings.Trim(v_string_six, v_cutset_no_trim), len(strings.Trim(v_string_six, v_cutset_no_trim)))                         //ikHo 4
	fmt.Println(strings.Trim(v_string_six, v_cutset_left_trim), len(strings.Trim(v_string_six, v_cutset_left_trim)))                     //ikHo 4
	fmt.Println(strings.Trim(v_string_six, v_cutset_right_trim), len(strings.Trim(v_string_six, v_cutset_right_trim)))                   //ikHo 4
	fmt.Println(strings.Trim(v_string_six, v_cutset_left_and_right_trim), len(strings.Trim(v_string_six, v_cutset_left_and_right_trim))) //ikHo 4

	//TrimLeft
	fmt.Println(strings.TrimLeft(v_string_seven, v_cutset_no_trim), len(strings.TrimLeft(v_string_seven, v_cutset_no_trim)))                         // TikHoT 7
	fmt.Println(strings.TrimLeft(v_string_seven, v_cutset_left_trim), len(strings.TrimLeft(v_string_seven, v_cutset_left_trim)))                     //ikHoT 5
	fmt.Println(strings.TrimLeft(v_string_seven, v_cutset_right_trim), len(strings.TrimLeft(v_string_seven, v_cutset_right_trim)))                   //ikHoT 5
	fmt.Println(strings.TrimLeft(v_string_seven, v_cutset_left_and_right_trim), len(strings.TrimLeft(v_string_seven, v_cutset_left_and_right_trim))) //ikHoT 5

	//TrimRight
	fmt.Println(strings.TrimRight(v_string_eight, v_cutset_no_trim), len(strings.TrimRight(v_string_eight, v_cutset_no_trim)))                         //TikHoT  7
	fmt.Println(strings.TrimRight(v_string_eight, v_cutset_left_trim), len(strings.TrimRight(v_string_eight, v_cutset_left_trim)))                     //TikHo 5
	fmt.Println(strings.TrimRight(v_string_eight, v_cutset_right_trim), len(strings.TrimRight(v_string_eight, v_cutset_right_trim)))                   //TikHo 5
	fmt.Println(strings.TrimRight(v_string_eight, v_cutset_left_and_right_trim), len(strings.TrimRight(v_string_eight, v_cutset_left_and_right_trim))) //TikHo 5

	fmt.Println(strings.Repeat("-", 20))
	//Trim
	fmt.Println(strings.Trim(v_string_nine, v_cutset_no_trim), len(strings.Trim(v_string_nine, v_cutset_no_trim)))                         // TikHoT  8
	fmt.Println(strings.Trim(v_string_nine, v_cutset_left_trim), len(strings.Trim(v_string_nine, v_cutset_left_trim)))                     //ikHo 4
	fmt.Println(strings.Trim(v_string_nine, v_cutset_right_trim), len(strings.Trim(v_string_nine, v_cutset_right_trim)))                   //ikHo 4
	fmt.Println(strings.Trim(v_string_nine, v_cutset_left_and_right_trim), len(strings.Trim(v_string_nine, v_cutset_left_and_right_trim))) //ikHo 4

	//TrimLeft
	fmt.Println(strings.TrimLeft(v_string_nine, v_cutset_no_trim), len(strings.TrimLeft(v_string_nine, v_cutset_no_trim)))                         // TikHoT  8
	fmt.Println(strings.TrimLeft(v_string_nine, v_cutset_left_trim), len(strings.TrimLeft(v_string_nine, v_cutset_left_trim)))                     //ikHoT  6
	fmt.Println(strings.TrimLeft(v_string_nine, v_cutset_right_trim), len(strings.TrimLeft(v_string_nine, v_cutset_right_trim)))                   //ikHoT  6
	fmt.Println(strings.TrimLeft(v_string_nine, v_cutset_left_and_right_trim), len(strings.TrimLeft(v_string_nine, v_cutset_left_and_right_trim))) //ikHoT  6

	//TrimRight
	fmt.Println(strings.TrimRight(v_string_nine, v_cutset_no_trim), len(strings.TrimRight(v_string_nine, v_cutset_no_trim)))                         // TikHoT  8
	fmt.Println(strings.TrimRight(v_string_nine, v_cutset_left_trim), len(strings.TrimRight(v_string_nine, v_cutset_left_trim)))                     // TikHo 6
	fmt.Println(strings.TrimRight(v_string_nine, v_cutset_right_trim), len(strings.TrimRight(v_string_nine, v_cutset_right_trim)))                   // TikHo 6
	fmt.Println(strings.TrimRight(v_string_nine, v_cutset_left_and_right_trim), len(strings.TrimRight(v_string_nine, v_cutset_left_and_right_trim))) // TikHo 6

	fmt.Printf("\n")

}

func variable_declaration_array() {

	fmt.Printf("variable_declaration_array\n")
	var v_array_one [2]int //Declare before assignment
	v_array_one[0] = 10
	v_array_one[1] = 20

	v_array_two := [5]int{10, 20, 30, 40, 50}     // Intialized with values
	var v_array_three [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println("&v_array_one:", &v_array_one, "&v_array_one[0]:", &v_array_one[0]) //&v_array_one: &[10 20] &v_array_one[0]: 0xc0000b2140
	fmt.Println(v_array_one)                                                        //[10 20]
	fmt.Println(v_array_two)                                                        //[10 20 30 40 50]
	fmt.Println(v_array_three)                                                      //[10 20 30 0 0]

	fmt.Printf("\n")
}

func variable_declaration_slice() {

	fmt.Printf("variable_declaration_slice\n")
	var v_slice_one []int //Declare before assignment
	v_slice_one = append(v_slice_one, 10, 20, 30, 40, 50)

	var v_slice_two []int = make([]int, 4)      //Len:4,Cap=4,auto resize 'Cap'
	var v_slice_three []int = make([]int, 2, 3) //Len:2,Cap=3
	v_slice_two[0] = 10
	v_slice_two[2] = 20
	v_slice_two[3] = 30
	//v_slice_two[4] = 40 //Index out of range [3] with length 3
	v_slice_two = append(v_slice_two, 40) //Use this

	v_slice_three[0] = 10
	v_slice_three[1] = 20

	var v_slice_four []int = []int{10, 20, 30, 40, 50} // Intialized with values

	fmt.Println("&v_slice_one:", &v_slice_one, "&v_slice_one[0]:", &v_slice_one[0]) //&v_slice_one: &[10 20 30 40 50] &v_slice_one[0]: 0xc0000b4090
	fmt.Println(v_slice_one)                                                        //[10 20 30 40 50]
	fmt.Println(v_slice_four)                                                       //[10 20 30 40 50]
	fmt.Println(v_slice_two, v_slice_three,
		"v_slice_two Len:", len(v_slice_two),
		"v_slice_two Cap:", cap(v_slice_two),
		"v_slice_three Len:", len(v_slice_three),
		"v_slice_three Cap:", cap(v_slice_three)) //[10 0 20 30 40] [10 20] v_slice_two Len: 5 v_slice_two Cap: 8 v_slice_three Len: 2 v_slice_three Cap: 3

	fmt.Printf("\n")

}

func variable_declaration_map() {
	fmt.Printf("variable_declaration_slice\n")
	var v_map_one = map[string]int{} //Empty Map,Declare before assignment
	v_map_one["key1"] = 1
	v_map_one["key2"] = 2
	v_map_one["key3"] = 3
	v_map_one["key4"] = 4

	var v_map_two = map[string]int{"key1": 10, "key2": 10, "key3": 10}

	v_map_three := make(map[string]float32)
	v_map_three["key1"] = 10.01
	v_map_three["key2"] = 20.01
	v_map_three["key3"] = 30.01

	fmt.Println(v_map_one, "v_map_one[\"key1\"]:", v_map_one[`key1`]) //map[key1:1 key2:2 key3:3 key4:4] v_map_one["key1"]: 1
	fmt.Println(v_map_two)                                            //map[key1:10 key2:10 key3:10]
	fmt.Println(v_map_three)                                          //map[key1:10.01 key2:20.01 key3:30.01]

	fmt.Printf("\n")
}

func variable_operators_builtin() {
	fmt.Printf("variable_operators_builtin\n")

	fmt.Printf("\n")
}

func variable_operators_collection() {
	//Add
	//Update
	//Delete
	//Get
	//Iterate
	//Sort
	//Merge

	fmt.Printf("variable_operators_collection\n")

	fmt.Printf("\n")
}
