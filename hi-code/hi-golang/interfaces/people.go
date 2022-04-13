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

type People struct {
}

func (p *People) BuyCarBySpeed(speed uint16) {
	switch {
	case speed >= 255:

		cars := []Car{new(FerrariCar)}
		for _, car := range cars {
			car.Run()
			fmt.Println(car.GetBrandName())
			car.UpdateColor("Yellow")
			car.Run()

		}
		//not break key
	case speed <= 140:
		var car Car
		car = &VolvoCar{} //car = VolvoCar{}
		car.Run()
		fmt.Println(car.GetBrandName())
		car.UpdateColor("White")
		car.Run()

	default:
		fmt.Print("No Selected Car")
	}
}
