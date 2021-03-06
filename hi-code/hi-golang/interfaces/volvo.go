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

type VolvoCar struct {
	speed     uint16
	brandName string
	color     string
}

func (v *VolvoCar) NewVolvoCar() *VolvoCar {
	return &VolvoCar{
		speed:     140,
		brandName: "Volvo",
		color:     "Black",
	}
}

func (v *VolvoCar) Run() {
	fmt.Println("This Car is ", v.brandName, ",Max speed is ", v.speed, "is color ", v.color)
}

func (f *VolvoCar) GetBrandName() string {
	return f.brandName
}

func (f *VolvoCar) UpdateColor(color string) {
	f.color = color
}
