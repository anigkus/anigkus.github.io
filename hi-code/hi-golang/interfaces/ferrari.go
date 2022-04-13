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

package interfaces

import "fmt"

type FerrariCar struct {
	speed     uint16
	brandName string
	color     string
}

func (f *FerrariCar) NewFerrariCar() *FerrariCar {
	return &FerrariCar{
		speed:     255,
		brandName: "Ferrari",
		color:     "Red",
	}
}

func (f *FerrariCar) Run() {

	fmt.Println("This Car is ", f.brandName, ",Max speed is ", f.speed, "is color ", f.color)
}

func (f *FerrariCar) GetBrandName() string {
	return f.brandName
}

func (f *FerrariCar) UpdateColor(color string) {
	f.color = color
}
