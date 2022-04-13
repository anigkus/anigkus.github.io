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
package appte

import (
	"fmt"
	"strings"

	"github.com/anigkus/hi-golang/constant"
	"github.com/anigkus/hi-golang/interfaces"
	"github.com/anigkus/hi-golang/structs"
	"github.com/anigkus/hi-golang/types"
	"github.com/anigkus/hi-golang/util"
)

func Main() {
	//App Banner
	text := fmt.Sprintln("Hi golang.")
	fmt.Println(text)

	//types package
	types.Main()

	//external needs to be new instance
	new(structs.Animal).Main()

	//external not to be new instance
	structs.NewAnimalNoParameter()
	structs.NewAnimalNoParameter().ToString()
	structs.NewAnimalNoParameter().AnimalNoParameter()
	structs.NewAnimalNoParameter().ToString()
	fmt.Println(structs.NewAnimalNoParameter().GetName())
	structs.NewAnimalNoParameter().SetName("xxx")
	fmt.Println(structs.NewAnimalNoParameter().GetName())
	structs.NewAnimalWithParamter("Fish", "Fishes")
	structs.NewAnimalNoParameter().ToString()

	//interface
	fmt.Println(strings.Repeat("-", 20))
	new(interfaces.People).BuyCarBySpeed(300)
	fmt.Println(strings.Repeat("-", 20))
	new(interfaces.People).BuyCarBySpeed(100)

	//constants
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(constant.BLACK)
	fmt.Println(constant.WHITE)
	fmt.Println(constant.RED)
	fmt.Println(constant.HOUR_ONE)
	fmt.Println(constant.ANIMAL_CAT)

	//util
	fmt.Println(strings.Repeat("-", 20))
	util.Main()
}
