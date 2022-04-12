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

package structs

import (
	"fmt"
	"strconv"
)

type Animal struct {
	name     string
	classify string
}

// instance function: public and not instance function
func (a *Animal) Main() {

	defaultAnimal := new(Animal).DefaultAnimal()
	fmt.Println("default Animal instance constructor", defaultAnimal.ToString())
	defaultAnimal.AnimalNoParameter()
	fmt.Println("no parameter constructor", defaultAnimal.ToString())

	//call internal fun
	internalOne()
	v, _ := internalFour("50")
	fmt.Println(internalTwo(), internalThree("30"), v)

}

// no parameter constructor: public and not instance function
func NewAnimalNoParameter() *Animal {
	return &Animal{
		name:     "People",
		classify: "Human",
	}
}

// with parameter constructor: public and not instance function
func NewAnimalWithParamter(name string, classify string) *Animal {

	return &Animal{
		name:     name,
		classify: classify,
	}
}

//public function ,no paramter and no return
func PublicOne() {
	//TODO
}

//public function ,no paramter and return one string type field
func PublicTwo() string {
	return "PublicFunTwo"
}

//public function, and return one int type field
func PublicThree(score string) int {
	v, _ := strconv.ParseInt(score, 0, 0)
	return int(v)
}

//public function , return two field and data type is (int,error)
func PublicFour(score string) (int, error) {
	v, err := strconv.ParseInt(score, 0, 0)
	return int(v), err
}

//Getter: public and instanced function
func (a *Animal) GetName() string {
	return a.name
}

//Getter: public and instanced function
func (a *Animal) GetClassify() string {
	return a.classify
}

//Setter: public and instanced function
func (a *Animal) SetName(name string) {
	a.name = name
}

//Setter: public and instanced function
func (a *Animal) SetClassify(classify string) {
	a.classify = classify
}

//default Animal instance: public and instanced function
func (a *Animal) DefaultAnimal() Animal {
	animal := Animal{
		name:     "People",
		classify: "Human",
	}
	return animal
}

// no parameter: public and instanced function
func (a *Animal) AnimalNoParameter() {
	a.name = "Dog"
	a.classify = "Canines"

}

// with parameter: public and instanced function
func (a *Animal) AnimalWithParameter(name string, classify string) {
	a.name = name
	a.classify = classify
}

//public and instanced function
func (a *Animal) ToString() string {
	return fmt.Sprintf("name:%v,classify:%v", a.name, a.classify)
}

//internal function ,no paramter and no return
func internalOne() {
	//TODO
}

//internal function ,no paramter and return one string type field
func internalTwo() string {
	return "internalTwo"
}

//internal function, and return one int type field
func internalThree(score string) int {
	v, _ := strconv.ParseInt(score, 0, 0)
	return int(v)
}

//internal function , return two field and data type is (int,error)
func internalFour(score string) (int, error) {
	v, err := strconv.ParseInt(score, 0, 0)
	return int(v), err
}
