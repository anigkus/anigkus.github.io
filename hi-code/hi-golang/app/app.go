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
	"github.com/anigkus/hi-golang/operator"
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
	animal := structs.NewAnimalNoParameter()
	fmt.Println(animal.ToString()) //name:People,classify:Human

	animal.AnimalNoParameter()
	animal.SetName("yyy")
	fmt.Println(animal.ToString()) //name:yyy,classify:Canines

	animal.AnimalNoParameter()
	fmt.Println(animal.ToString()) //name:Dog,classify:Canines

	fmt.Println(structs.NewAnimalNoParameter().ToString()) //name:People,classify:Human
	fmt.Println(structs.NewAnimalNoParameter().GetName())  //People

	animal = structs.NewAnimalNoParameter()
	animal.SetName("xxx")
	fmt.Println(animal.GetName()) //xxx

	fmt.Println(structs.NewAnimalNoParameter().GetName())        //People
	fmt.Println(structs.NewAnimalWithParamter("Fish", "Fishes")) //&{Fish Fishes}
	fmt.Println(structs.NewAnimalNoParameter().ToString())       //name:People,classify:Human

	//interface
	fmt.Println(strings.Repeat("-", 20))
	new(interfaces.People).BuyCarBySpeed(300)
	fmt.Println(strings.Repeat("-", 20))
	new(interfaces.People).BuyCarBySpeed(100)

	//constants
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(constant.BLACK)      //#000000
	fmt.Println(constant.WHITE)      //#FFFFFF
	fmt.Println(constant.RED)        //#FF0000
	fmt.Println(constant.HOUR_ONE)   //1
	fmt.Println(constant.ANIMAL_CAT) //CAT

	//util
	fmt.Println(strings.Repeat("-", 20))
	util.Main()

	//operator
	operator.Main()
}
