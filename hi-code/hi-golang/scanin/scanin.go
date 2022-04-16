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

package scanin

import (
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	fmt.Println("scan...")
	//
	buyGoods()
}

func buyGoods() {

	sortIndex := 0
	buy := true
	continueBuy := ""
	fmt.Println("Welcom goods store.")
	var goodLists []map[string]string // unordered collection
	for buy {
		goods := scanlnGoods(strconv.Itoa(sortIndex))
		goodLists = append(goodLists, goods)
		fmt.Print("Do you continue buy? Y/N: ")
		fmt.Scanln(&continueBuy)
		continueBuy = strings.ToUpper(continueBuy)
		if continueBuy == "N" {
			break
		} else if continueBuy == "" {
			panic(fmt.Sprintf("You Input Error:%v", continueBuy))
		} else {
			continue
		}
	}
	//sort
	if len(goodLists) != 0 {
		fmt.Println("You buy goods list:")
		var sumGoodsPrice float64 = 0
		var sumGoodsNumber int = 0

		for index, goodList := range goodLists {
			goodsInfo := ""
			for key, value := range goodList {
				goodsInfo += value + "\t"
				if key == "goodsPriceA" {
					v, _ := strconv.ParseFloat(value, 64)
					sumGoodsPrice += v
				}
				if key == "goodsNumberA" {
					v, _ := strconv.Atoi(value)
					sumGoodsNumber += v
				}
			}
			number := index + 1
			fmt.Printf("%v. %v\n", number, goodsInfo)
		}
		fmt.Printf("Good Sum Number:%v. Good Sum Price:%v\n", sumGoodsNumber, sumGoodsPrice)

	} else {
		fmt.Println("You not buyed goods.")
	}

}

func scanlnGoods(sortIndex string) (goods map[string]string) {
	var goodsName string
	var goodsPrice string
	var goodsNumber string
	fmt.Print("Please input name:")
	fmt.Scanln(&goodsName)
	fmt.Print("Please input price:")
	fmt.Scanln(&goodsPrice)
	fmt.Print("Please input number:")
	fmt.Scanln(&goodsNumber)
	goods = make(map[string]string)
	goods["sortIndex"] = sortIndex
	goods["goodsNameA"] = goodsName
	goods["goodsPriceB"] = goodsPrice
	goods["goodsNumberC"] = goodsNumber
	return goods
}
