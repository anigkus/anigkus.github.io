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

package conditional

import (
	"fmt"
	"strconv"
	"time"
)

func Main() {
	//if
	//if...else
	//if... else if...else
	//switch...case
	fmt.Println("conditional...")
	ifStatement()

	switchStatement()

	selectStatement()
}

func ifStatement() {
	int64_one := 51
	if int64_one > 100 {
		fmt.Println("int64_one>100 is true")
	}
	if int64_one > 100 {
		fmt.Println("int64_one>100 is true")
	} else {
		fmt.Println("int64_one>100 is not true") //
	}
	if int64_one > 100 {
		fmt.Println("int64_one>100 is true")
	} else if int64_one > 50 {
		fmt.Println("int64_one>50 is true") //
	} else {
		fmt.Println("int64_one else:", int64_one)
	}
}

func switchStatement() {
	var int64_one int64 = 51
	switch int64_one {
	case 50:
		fmt.Println("case:", int64_one)
		//break
		//not explicitly break  ,With break by default
	case 51:
		fmt.Println("case:51") //case:51
	case 52:
		fmt.Println("case:52")
	case 53:
		fmt.Println("case:53")
	default:
		fmt.Println("default")
	}
	switch int64_one {
	case 50:
		fmt.Println("case 50")
	case 51:
		fmt.Println("case:51") //case:51
		fallthrough
	case 52:
		fmt.Println("case:52") //case:52
	case 53:
		fmt.Println("case:53")
	default:
		fmt.Println("default")
	}
	switch {
	case int64_one == 50:
		fmt.Println("case 50")
	case int64_one == 51:
		fmt.Println("case int64_one == 51") //cae int64_one == 51
	case int64_one == 52:
		fmt.Println("case:52")
	case int64_one == 53:
		fmt.Println("case:53")
	default:
		fmt.Println("default")
	}

	switch int64_one {
	case 10, 20:
		fmt.Println("case 10, 20")
	case 30, 40:
		fmt.Println("case:30, 40")
	case 50, 51, 52, 60:
		fmt.Println("case:50, 51, 52, 60") //case:50, 51, 52, 60
	case 70, 80, 90, 100:
		fmt.Println("case:70, 80, 90, 100")
	default:
		fmt.Println("default")
	}

	switch strconv.FormatInt(int64_one, 10) {
	case "50":
		fmt.Println("case 50")
	case "51":
		fmt.Println("case:51") //case:51
	case "52":
		fmt.Println("case:52")
	case "53":
		fmt.Println("case:53")
	default:
		fmt.Println("default")
	}

	whatType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Is bool")
		case int:
			fmt.Println("Is int")
		default:
			fmt.Println("Unknow", t)
		}
	}
	whatType(true) //Is bool
	whatType(1)    //Is int
	whatType('c')  //Unknow 99
}

func selectStatement() {
	char_one := make(chan int)
	char_two := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		char_one <- 1
	}()
	go func() {
		time.Sleep(time.Second * 1)
		char_two <- 2
	}()
	index := 2
	for index > 0 {
		select {
		case msg_one := <-char_one:
			fmt.Println("received:", msg_one) //received: 1
		case msg_two := <-char_two:
			fmt.Println("received:", msg_two) //received: 2
		}
		index--
	}
	close(char_one) // Closes the channel
	close(char_two) // Closes the channel

}
